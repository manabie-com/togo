package utils

import "time"

var DEFAULT_FORMAT_DATE = "2006-01-02"

func FormatTimeToString(time time.Time) string {
	return time.Format(DEFAULT_FORMAT_DATE)
}
