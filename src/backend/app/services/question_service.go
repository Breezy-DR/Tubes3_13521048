package services

import (
	"backend/app/gorms"
	"backend/app/models"
	"backend/internal/algorithm"
	"backend/internal/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type QuestionService struct {
	g gorms.Gorms
}

type SimilarityPair struct {
	similarity float64
	question   models.QuestionNA
}

const CalculatorRegex string = "\\s*(([\\+\\-\\*\\/()^]|\\d)+\\s*)*\\s*"
const DateRegex string = "((\\d{4}[\\/-]\\d{1,2}[\\/-]\\d{1,2})|(\\d{1,2}[\\/-]\\d{1,2}[\\/-]\\d{4}))"
const AddRegex string = "(?i)tambahkan pertanyaan\\s+[\\x00-\\x7F]+[\\x00-\\x7F\\s]*\\s+dengan\\s+jawaban\\s+[\\x00-\\x7F]+[\\x00-\\x7F\\s]*"
const RemoveRegex string = "(?i)\\s*hapus\\s+pertanyaan\\s+[\\x00-\\x7F\\s]+[\\x00-\\x7F\\s\\s]*"

const notFound = "pertanyaan tidak ditemukan di database"
const unknownErr = "unknown error has occured"

func (q QuestionService) GetAnswer(question string, sessionID string, searchAlg string) (string, bool) {
	found := false
	res := ""
	if reg, _ := regexp.Compile(DateRegex); util.StringMatches(question, reg) {
		res = algorithm.GetDayDate(question)
		found = true
	} else if reg, _ := regexp.Compile(CalculatorRegex); !found && util.StringMatches(question, reg) {
		// TODO: iterate regex instead of this
		val, err := algorithm.CalculateExpression(question)

		if err == nil {
			res = strconv.FormatFloat(val, 'f', 6, 64)
			found = true
		}
	} else if reg, _ := regexp.Compile(AddRegex); !found && util.StringMatches(question, reg) {
		entries := strings.Split(question, " ")
		// skip: "tambahkan pertanyaan"
		i := 2

		tReg, _ := regexp.Compile("(?i)dengan")

		for i < len(entries) && !util.StringMatches(entries[i], tReg) {
			i++
		}

		if i >= len(entries) {
			return unknownErr, false
		}
		nQuestion := strings.Join(entries[2:i], " ")

		// skip: "dengan jawaban"
		i += 2
		if i >= len(entries) {
			return unknownErr, false
		}
		ans := strings.Join(entries[i:], " ")

		fmt.Println(nQuestion)
		fmt.Println(ans)

		q.g.Gorm.Create(&models.QuestionNA{Question: nQuestion, Answer: ans})

		found = true
		res = "Pertanyaan berhasil ditambahkan"
	} else if reg, _ := regexp.Compile(RemoveRegex); !found && util.StringMatches(question, reg) {
		//extractorRegex := "\\b(?!tambahkan\\b)(?!pertanyaan\\b)(?!dengan\\b)(?!jawaban\\b)\\w+\\s*"
		entry := strings.SplitN(question, " ", 3)[2]
		ans, ansArr := q.findMatch(entry, searchAlg)

		found = true
		if ansArr == nil {
			q.g.Gorm.Delete(ans)
			res = "Pertanyaan berhasil dihapus"
		}
		res = notFound
	} else if !found {
		// find with string matcher
		ans, ansArr := q.findMatch(question, searchAlg)

		found = true
		if ansArr == nil {
			res = ans.Answer
		} else {
			var sb strings.Builder
			sb.WriteString(notFound + "\n" + "Apakah maksud anda:\n")
			for i := 0; i < len(ansArr); i++ {
				sb.WriteString(ansArr[i].Question)
				sb.WriteString("\n")
			}
			res = sb.String()
		}
	}

	q.g.Gorm.Create(&models.History{SessionID: sessionID, UserEntry: question, Answer: res, CreationDate: time.Now()})

	return res, found
}

func (q QuestionService) findMatch(search, searchAlg string) (models.QuestionNA, []models.QuestionNA) {
	var (
		questions  []models.QuestionNA
		topMatches []SimilarityPair
		found      = false
		first      = true
		result     models.QuestionNA
		idStart    = 0
	)

	for !found && (first || len(questions) > 0) {
		q.g.Gorm.Where("question_id BETWEEN ? AND ?", idStart, idStart+25).Find(&questions)
		for i := 0; !found && i < len(questions); i++ {
			question := questions[i].Question

			idx := algorithm.SearchKMP(search, question)

			if idx != -1 {
				found = true
				result = questions[i]
			} else {
				distance := algorithm.GetLevDistance(search, question)
				similarity := float64(len(search)-distance) / float64(len(question))
				addToSimPair(similarity, questions[i], &topMatches)
			}

		}
		first = false
		idStart += 25
	}

	if found {
		return result, nil
	}

	if maxIdx := getMaxSimilarity(topMatches); maxIdx != -1 && topMatches[maxIdx].similarity >= 0.9 {
		return topMatches[maxIdx].question, nil
	}

	var resArr = make([]models.QuestionNA, len(topMatches))

	for i := 0; i < len(topMatches); i++ {
		resArr[i] = topMatches[i].question
	}

	return models.QuestionNA{}, resArr
}

func addToSimPair(similarity float64, question models.QuestionNA, arr *[]SimilarityPair) {
	if len(*arr) < 3 {
		*arr = append(*arr, SimilarityPair{similarity: similarity, question: question})
		return
	}

	minIdx := 0
	var minSim = (*arr)[0].similarity

	for i := 1; i < len(*arr); i++ {
		if (*arr)[i].similarity < minSim {
			minSim = (*arr)[i].similarity
			minIdx = i
		}
	}

	if minSim > similarity {
		return
	}
	(*arr)[minIdx] = SimilarityPair{similarity: similarity, question: question}
}

func getMaxSimilarity(arr []SimilarityPair) int {
	if len(arr) == 0 {
		return -1
	}

	minIdx := 0
	var minSim = arr[0].similarity
	for i := 1; i < 3; i++ {
		if arr[i].similarity < minSim {
			minSim = arr[i].similarity
			minIdx = i
		}
	}

	return minIdx
}

func NewQuestionService(g gorms.Gorms) QuestionService {
	return QuestionService{g: g}
}
