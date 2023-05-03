package algorithm

import (
	"backend/internal/util"
	"math"
)

func GetLevDistance(searchStr, text string) int {
	var (
		textLen    = len(text)
		searchLen  = len(searchStr)
		distMatrix = make([][]int, textLen+1)
	)

	//fmt.Println(textLen, searchLen)

	distMatrix[0] = make([]int, searchLen+1)
	for i := 0; i < searchLen+1; i++ {
		distMatrix[0][i] = i
	}

	for i := 1; i < textLen+1; i++ {
		distMatrix[i] = make([]int, searchLen+1)
		distMatrix[i][0] = i

	}

	for i := 1; i < textLen+1; i++ {
		for j := 1; j < searchLen+1; j++ {
			temp := int(math.Min(math.Min(float64(distMatrix[i-1][j-1]), float64(distMatrix[i][j-1])), float64(distMatrix[i-1][j])))
			if !util.EqualCaseIns(text[i-1], searchStr[j-1]) {
				temp += 1
			}
			distMatrix[i][j] = temp
		}
	}
	//for i := 0; i < textLen+1; i++ {
	//	fmt.Println(distMatrix[i])
	//}
	return distMatrix[textLen][searchLen]
}
