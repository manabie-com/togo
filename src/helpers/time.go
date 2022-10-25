package helpers

import "time"

func BeginOfDate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.UTC().Location())
}

func EndOfDate(t time.Time) time.Time {
	oneDay := 24 * time.Hour
	return BeginOfDate(t).Add(oneDay)
}
