package tools

import (
	"fmt"
	"time"
)

func GetTimeDayString() string {

	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	DayString := fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	return DayString
}
