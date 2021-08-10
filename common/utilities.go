package common

import "time"

func GetCurrentDateRounded() time.Time {
	now := time.Now()
	nowRounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return nowRounded
}
