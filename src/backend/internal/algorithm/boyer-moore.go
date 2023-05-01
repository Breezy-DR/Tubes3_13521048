package algorithm

func SolveBm(pattern, text string) {

	//charShiftTable := getCharShiftTable(pattern)
	//i := 0
	//j := 0
	//var textLen = len(text)
	//var match = false
	//for i < textLen && !match {
	//
	//}

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
