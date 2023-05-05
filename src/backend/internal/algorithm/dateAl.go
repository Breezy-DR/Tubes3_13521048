package algorithm

import (
	"backend/internal/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetDayDate(dateStr string) string {

	now := "0001-01-01"
	num := int(math.Mod(math.Abs(float64(getX(dateStr)-getX(now))), 7))
	var day = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	return day[num]

}

func getX(dateStr string) int {
	var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var year float64
	var month float64
	var day float64
	fmt.Println(dateStr)
	if util.StringMatches(dateStr, DateFirstRe) {
		var dateList []string
		if dateStr[4] == '/' {
			dateList = strings.Split(dateStr, "/")
		} else {
			dateList = strings.Split(dateStr, "-")
		}

		year, _ = strconv.ParseFloat(dateList[0], 64)
		month, _ = strconv.ParseFloat(dateList[1], 64)
		day, _ = strconv.ParseFloat(dateList[2], 64)
	} else if util.StringMatches(dateStr, DateSecondRe) {
		var dateList []string
		if dateStr[len(dateStr)-5] == '/' {
			dateList = strings.Split(dateStr, "/")
		} else {
			dateList = strings.Split(dateStr, "-")
			// extra safe
			if len(dateList) == 1 {
				dateList = strings.Split(dateStr, "/")
			}
		}
		fmt.Println(dateList)
		year, _ = strconv.ParseFloat(dateList[2], 64)
		month, _ = strconv.ParseFloat(dateList[1], 64)
		day, _ = strconv.ParseFloat(dateList[0], 64)
	}

	dayInY := int(year)*365 + int(math.Floor(year/4.0)) - int(math.Floor(year/100)) + int(math.Floor(year/400))
	dayInM := 0
	for i := 0; i < int(month)-1; i++ {
		dayInM += days[i]
	}
	dayInD := day
	fmt.Println(dayInY, dayInM, dayInD)

	return dayInY + dayInM + int(dayInD)
}
