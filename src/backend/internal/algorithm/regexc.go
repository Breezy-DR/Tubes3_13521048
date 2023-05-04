package algorithm

import "regexp"

var MultiQuestionsRe = regexp.MustCompile("[?;]")
var CalculatorRe = regexp.MustCompile(CalculatorRegexString)
var DateRe = regexp.MustCompile(DateRegexString)
var AddRe = regexp.MustCompile(AddRegexString)
var RemoveRe = regexp.MustCompile(RemoveRegexString)

var DateFirstRe = regexp.MustCompile(dates[0])
var DateSecondRe = regexp.MustCompile(dates[1])

var DenganRe = regexp.MustCompile("(?i)dengan")
var JawabanRe = regexp.MustCompile("(?i)jawaban")
