package timestamp

import "time"

const (
	dateFormat = "2006-01-02"
)

func GetCurrentTime() string {
	now := time.Now()

	return now.Format(dateFormat)
}
