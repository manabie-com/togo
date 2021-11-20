package model

import (
	"math"
	"time"
)

func getDateTimestamp(unixTime float64) time.Time {
	de, s := math.Modf(unixTime)
	return time.Unix(int64(de), int64(s*1e9)).UTC()
}

func getFlowTimestamp(milisecondTime uint64) string {
	unixTime := float64(milisecondTime / 1000000)
	dateTime := getDateTimestamp(unixTime)
	return dateTime.Format(time.RFC3339Nano)
}
