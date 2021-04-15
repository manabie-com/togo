package util

import (
	"errors"
	"time"
)

func ValidateCreatedDate(createdDate string) error {
	if createdDate == "" {
		return errors.New("created_date is missing")
	}
	_, err := time.Parse("2006-01-02", createdDate)
	if err != nil {
		return errors.New("created_date is invalid")
	}
	return nil
}
