package algorithm

import (
	"math"
)

func SearchBM(pattern, text string) int {

	charShiftTable := getCharShiftTable(pattern)
	var (
		textLen    = len(text)
		patternLen = len(pattern)
		j          = patternLen - 1
		idx        = 0
	)

	if patternLen > textLen {
		return -1
	}
	for j > -1 && idx < textLen {
		if text[idx] != pattern[j] {
			ni := charShiftTable[text[idx]]
			idx += patternLen - int(math.Min(float64(j), float64(ni+1)))
			j = patternLen - 1

			continue
		}
		idx--
		j--
	}

	if j != -1 {
		return -1
	}

	return idx + 1
}

func getCharShiftTable(pattern string) map[uint8]int {
	var (
		patternLen = len(pattern)
		shiftMap   = make(map[uint8]int)
	)

	for i := 0; i < patternLen; i++ {
		shiftMap[pattern[i]] = i
	}

	return shiftMap
}
