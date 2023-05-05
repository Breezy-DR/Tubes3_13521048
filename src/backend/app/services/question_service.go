package services

import (
	"backend/app/gorms"
	"backend/app/models"
	"backend/internal/algorithm"
	"backend/internal/util"
	"fmt"
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

const notFound = "pertanyaan tidak ditemukan di database."
const unknownErr = "unknown error has occured"

func (q QuestionService) GetAnswer(question string, sessionID string, searchAlg string) ([]string, bool) {
	found := false
	questions := algorithm.MultiQuestionsRe.Split(question, -1)
	var histories = make([]models.History, len(questions))
	var responses = make([]string, len(questions))
	timeNow := time.Now()

	j := 0
	for i := 0; i < len(questions); i++ {
		currentQuestion := strings.Trim(questions[i], " ")

		if currentQuestion == "" {
			responses = responses[:len(responses)-1]
			histories = histories[:len(histories)-1]
			continue
		}

		temp, f := q.getResponse(currentQuestion, searchAlg)
		found = found || f
		responses[j] = temp
		histories[j] = models.History{SessionID: sessionID, UserEntry: currentQuestion, Answer: temp, CreationDate: timeNow}
		j++
	}

	q.g.Gorm.Create(&histories)

	return responses, found
}

func (q QuestionService) getResponse(question string, searchAlg string) (string, bool) {
	found := false
	res := ""

	if util.StringMatches(question, algorithm.DateRe) {
		res = algorithm.GetDayDate(question)

		found = true
	} else if !found && util.StringMatches(question, algorithm.CalculatorRe) {
		val, err := algorithm.CalculateExpression(question)

		if err == nil {
			res = strconv.FormatFloat(val, 'f', 6, 64)
			found = true
		}
	} else if !found && util.StringMatches(question, algorithm.AddRe) {
		entries := strings.Split(question, " ")
		// skip: "tambahkan pertanyaan"
		i := 0
		j := 2

		for j < len(entries) {
			if util.StringMatches(entries[j], algorithm.DenganRe) &&
				j+1 < len(entries) && util.StringMatches(entries[j+1], algorithm.JawabanRe) {
				i = j
			}
			j++
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

		q.g.Gorm.Create(&models.QuestionNA{Question: nQuestion, Answer: ans})

		found = true
		res = "Pertanyaan berhasil ditambahkan"
	} else if !found && util.StringMatches(question, algorithm.RemoveRe) {
		entry := strings.SplitN(question, " ", 3)[2]
		ans, ansArr := q.findMatch(entry, searchAlg)

		found = true
		if ansArr == nil {
			q.g.Gorm.Delete(ans)
			res = "Pertanyaan berhasil dihapus"
		} else {
			res = notFound
		}
	} else if !found {
		// find with string matcher
		ans, ansArr := q.findMatch(question, searchAlg)

		found = true
		if ansArr == nil {
			res = ans.Answer
		} else {
			var sb strings.Builder
			sb.WriteString(notFound)
			if len(ansArr) > 0 {
				sb.WriteString(" Apakah maksud anda: ")
				for i := 0; i < len(ansArr); i++ {
					sb.WriteString(strconv.Itoa(i + 1))
					sb.WriteString(". ")
					sb.WriteString(ansArr[i].Question)
					sb.WriteString(" ")
				}
			}
			res = sb.String()
		}
	}
	return res, found
}

func (q QuestionService) findMatch(search, searchAlg string) (models.QuestionNA, []models.QuestionNA) {
	var (
		questions        []models.QuestionNA
		topMatches       []SimilarityPair
		found            = false
		foundExact       = false
		result           models.QuestionNA
		shortestMatchLen = -1
	)
	q.g.Gorm.Order("question_id asc").Limit(25).Find(&questions)

	for !foundExact && len(questions) > 0 {
		for i := 0; !foundExact && i < len(questions); i++ {
			question := questions[i].Question

			var idx int
			if searchAlg == algorithm.KMP {
				idx = algorithm.SearchKMP(search, question)
			} else {
				idx = algorithm.SearchBM(search, question)
			}

			if idx != -1 {
				found = true
				if len(search) == len(question) {
					foundExact = true
					result = questions[i]
				} else if shortestMatchLen == -1 || shortestMatchLen > len(question) {
					result = questions[i]
					shortestMatchLen = len(question)
				}
			} else {
				distance := algorithm.GetLevDistance(search, question)
				similarity := float64(len(question)-distance) / float64(len(question))
				addToSimPairs(similarity, questions[i], &topMatches)
			}

		}
		largestIndex := questions[len(questions)-1].QuestionID
		q.g.Gorm.Where("question_id > ?", largestIndex).Find(&questions)
	}

	if found {
		return result, nil
	}

	if maxIdx := getMaxSimilarity(topMatches); maxIdx != -1 && topMatches[maxIdx].similarity >= 0.9 {
		fmt.Println(topMatches)
		return topMatches[maxIdx].question, nil
	}

	var resArr = make([]models.QuestionNA, len(topMatches))

	for i := 0; i < len(topMatches); i++ {
		resArr[i] = topMatches[i].question
	}

	return models.QuestionNA{}, resArr
}

func addToSimPairs(similarity float64, question models.QuestionNA, arr *[]SimilarityPair) {
	if similarity < 0.7 {
		return
	}

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
	for i := 1; i < len(arr); i++ {
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
