package utils

import "time"

func EndOfCurrentDate(t time.Time) time.Time {
	nextDayStr := t.Format("2006-01-02") + "T" + "23:59:59"
	nextDay, _ := time.Parse("2006-01-02T15:04:05", nextDayStr)
	return nextDay
}
