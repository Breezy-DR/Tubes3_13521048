package algorithm

import (
	"backend/internal/util"
)

func SearchKMP(pattern, text string) int {
	var (
		patternLen = len(pattern)
		revIdx     = getJumpArr(pattern)
		textLen    = len(text)
	)

	if patternLen > textLen {
		return -1
	}

	var (
		j       = -1
		idx     = 0
		inMatch = false
		found   = false
	)
	for !found && idx < textLen {
		if inMatch && j < (patternLen-1) && util.EqualCaseIns(pattern[j+1], text[idx]) {
			j += 1
		} else if inMatch {
			foundRev := false
			for j != -1 && !foundRev {
				j = revIdx[j] - 1
				if j < (patternLen-1) && pattern[j+1] == text[idx] {
					foundRev = true
					j += 1
				}
			}
			if j == -1 {
				inMatch = false
			}
		} else if j < (patternLen-1) && util.EqualCaseIns(pattern[j+1], text[idx]) {
			inMatch = true
			j += 1

		}
		idx += 1

		if inMatch && j == patternLen-1 {
			found = true
		}
	}

	if found {
		return idx - patternLen
	}
	return -1
}

func getJumpArr(pattern string) []int {
	var (
		patternLen = len(pattern)
		revIdx     = make([]int, patternLen)
	)
	revIdx[0] = 0

	var (
		cont    = false
		contIdx = 0
	)
	for i := 1; i < patternLen; i++ {
		if cont && pattern[i] == pattern[contIdx] {
			revIdx[i] = contIdx + 1
			contIdx += 1
			continue
		} else if cont && pattern[i] != pattern[contIdx] {
			cont = false
			contIdx = 0
			revIdx[i] = 0
			continue
		}

		if pattern[i] == pattern[0] {
			cont = true
			revIdx[i] = 1
			contIdx = 1
		}
	}

	return revIdx
}
