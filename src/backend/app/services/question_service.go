package services

import (
	"backend/internal/algorithm"
	"backend/internal/util"
	"regexp"
	"strconv"
)

type QuestionService struct {
}

const CalculatorRegex string = "\\s*(([\\+\\-\\*\\/()^]|\\d)+\\s*)*\\s*"

const DateRegex string = "((\\d{4}[\\/-]\\d{1,2}[\\/-]\\d{1,2})|(\\d{1,2}[\\/-]\\d{1,2}[\\/-]\\d{4}))"

func (q QuestionService) GetAnswer(question string) string {

	if reg, _ := regexp.Compile(DateRegex); util.StringMatches(question, reg) {
		return algorithm.GetDayDate(question)
	} else if reg, _ := regexp.Compile(CalculatorRegex); util.StringMatches(question, reg) {
		res, err := algorithm.CalculateExpression(question)

		if err == nil {
			return strconv.FormatFloat(res, 'f', 6, 64)
		}

	}

	return "sd"
}

func NewQuestionService() QuestionService {
	return QuestionService{}
}
