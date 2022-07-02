package rest

import (
	"errors"
	"strings"
	"time"
)

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Validate the username if empty string or if it exceed the character limit
func validateUsername(username string) error {
	if trimLowerUsername(username) == "" {
		return errors.New("username is required")
	} else if len(trimLowerUsername(username)) > 200 {
		return errors.New("username has 200 character limit")
	}

	return nil
}

// Validate the taskDailyLimit less than 1
func validateTaskDailyLimit(limit int32) error {
	if limit < 1 {
		return errors.New("taskDailyLimit must be atleast 1")
	}
	return nil
}

// Validate the title if empty string or if it exceed the character limit
func validateTaskTitle(title string) error {
	if title == "" {
		return errors.New("username is required")
	} else if len(title) > 200 {
		return errors.New("title has 200 character limit")
	}
	return nil
}

// Trim and transform to lowercase the username
func trimLowerUsername(username string) string {
	return strings.TrimSpace(strings.ToLower(username))
}
