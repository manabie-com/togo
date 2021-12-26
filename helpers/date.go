package helpers

import "time"

func GetDateNow() time.Time {
	return time.Now().UTC()
}
