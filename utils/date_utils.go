package utils

import "time"

const (
	dateLayout  = "2006-01-02 15:04:05.000Z"
	todayLayout = "2006-01-02"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowFormat() string {
	return GetNow().Format(dateLayout)
}

func GetToday() string {
	return GetNow().Format(todayLayout)
}
