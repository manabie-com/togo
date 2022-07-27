package utils

import "time"

func StartOfDay(current time.Time) time.Time {
	y, m, d := current.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, current.Location())
}

func EndOfDay(current time.Time) time.Time {
	y, m, d := current.Date()
	return time.Date(y, m, d, 23, 59, 59, 0, current.Location())
}
