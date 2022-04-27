package common

import "time"

type Time struct {
	time.Time
}

func MakeTime(
	iTime time.Time,
) Time {
	return Time {
		Time: iTime,
	}
}