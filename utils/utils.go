package utils

import "time"

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("02-01-2006", stringDate)
	return err == nil
}
