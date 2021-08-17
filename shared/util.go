package shared

import (
	"fmt"
	"time"
)

func FormatDateString(now time.Time) string {
	year, month, day := now.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
