package algorithm

import (
	"backend/internal/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const DateRegexF string = "(\\d{4}[\\/-]\\d{1,2}[\\/-]\\d{1,2})"
const DateRegexS string = "(\\d{1,2}[\\/-]\\d{1,2}[\\/-]\\d{4})"

func GetDayDate(dateStr string) string {

	now := "2023-05-01"
	num := int(math.Mod(float64(getX(dateStr)-getX(now)), 7))
	var day = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	return day[num]

}

func getX(dateStr string) int {
	var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if reg, _ := regexp.Compile(DateRegexF); util.StringMatches(dateStr, reg) {
		var dateList []string
		if dateStr[4] == '/' {
			dateList = strings.Split(dateStr, "/")
		} else {
			dateList = strings.Split(dateStr, "-")
		}

		year, _ := strconv.ParseFloat(dateList[0], 64)
		month, _ := strconv.ParseFloat(dateList[1], 64)
		day, _ := strconv.ParseFloat(dateList[2], 64)

		dayInY := int(year)*365 + int(math.Floor(year/4.0)) - int(math.Floor(year/100)) + int(math.Floor(year/400))
		dayInM := 0
		for i := 0; i < int(month)-1; i++ {
			dayInM += days[i]
		}
		dayInD := day

		return dayInY + dayInM + int(dayInD)
	}

	return 0
}
