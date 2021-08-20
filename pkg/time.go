package pkg

import "time"

func UnixTimeUntilEndOfDay(t time.Time) int64 {
	y, m, d := t.Date()
	endOfDay := time.Date(y, m, d, 23, 59, 59, 0, t.Location())
	remaining := endOfDay.Unix() - t.Unix()
	return remaining
}
