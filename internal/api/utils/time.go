package utils

import "time"

const (
	DefaultLayout = "2006-01-02"
)

func GetTimeNowWithDefaultLayoutInString() string {
	return time.Now().Format(DefaultLayout)
}
