package utils

import "time"

var VNLocation, _ = time.LoadLocation("Asia/Ho_Chi_Minh")

func Now() time.Time {
	return time.Now().In(VNLocation)
}

func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

func NowInUnixSecond() int64 {
	return TimeToUnix(Now())
}
