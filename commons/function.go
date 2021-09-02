package commons

import (
	"iris-demo/config"
	"iris-demo/slog"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const TIME_DAY_LAYOUT = "02"

func SetLocation() *time.Location {
	local, err := time.LoadLocation(config.Setting.App.Timezone)
	if err == nil {
		return local
	}
	slog.Errorf(" timestamp err:%s", err.Error())
	return time.Local
}

// time to string time
func TimeToString(t *time.Time) string {
	return t.In(SetLocation()).Format(TIME_LAYOUT)
}

// get time corresponding day
func GetDay(t *time.Time) string {
	return t.In(SetLocation()).Format(TIME_DAY_LAYOUT)
}

// get current month's the last day
func GetCurrentMonthsLastDay() string {
	nowTime := time.Now()
	currentYear,currentMonth,_ := nowTime.Date()
	startTime := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, SetLocation())
	endTime := startTime.AddDate(0, 1, 0).Add(-time.Second)
	lastDay := GetDay(&endTime)
	return lastDay
}
