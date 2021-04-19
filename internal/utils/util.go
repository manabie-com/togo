package utils

import (
	"time"
)

func LessThanLimit(start_time string, limit_time time.Duration) bool {
	x1, _ := time.Parse(time.UnixDate, start_time)
	x2 := time.Now()

	if x2.Sub(x1) < limit_time {
		return true
	}

	return false
}

func GetStartTime() string {
	return time.Now().Format(time.UnixDate)
}