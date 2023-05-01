package util

import "regexp"

func StringMatches(text string, re *regexp.Regexp) bool {
	matches := re.FindAllString(text, -1)
	return len(matches) == 1 && matches[0] == text
}

func EqualCaseIns(a, b uint8) bool {
	if a == b {
		return true
	}

	if !isAlphabet(a) || !isAlphabet(b) {
		return false
	}

	if a > b && (a-32) == b {
		return true
	}

	if b > a && (b-32) == a {
		return true
	}

	return false
}

func isAlphabet(a uint8) bool {
	return (a >= 65 && a <= 90) || (a >= 97 && a <= 122)
}
