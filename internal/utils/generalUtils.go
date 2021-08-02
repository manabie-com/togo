package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 55 {
		return false
	}
	return emailRegex.MatchString(e)
}

// isEmailValid checks if the email provided passes the required structure and length.
func IsUsernameValid(e string) (bool, error) {
	if e != "" {
		if len(e) < 3 || len(e) > 55 {
			return false, errors.New("invalid username length (3 to 55 characters)")
		}
		return true, nil
	} else {
		return false, errors.New(`username contains ""`)
	}
}

// GetCurrentTime function is used to get the current time in milliseconds.
func GetCurrentEpochTimeInMiliseconds() int64 {
	var now = time.Now()
	ts := now.UnixNano() / 1000000
	return ts
}

func GetTimesByPeriod(period string) (time.Time, time.Time, error) {
	var startDate time.Time
	var endDate time.Time
	var timeNow = time.Now()
	switch period {

	case "daily":
		{
			startDate = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
			endDate = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, time.Local)
			break
		}
	case "monthly":
		{
			fmt.Println("================= monthly filter")
			startDate = time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.Local)
			endDate = time.Date(timeNow.Year(), timeNow.Month()+1, 0, 23, 59, 59, 0, time.Local)
			break
		}
	case "yearly":
		{
			fmt.Println("========= Yearly filter")
			startDate = time.Date(timeNow.Year(), 1, 1, 00, 00, 00, 00, time.Local)
			endDate = time.Date(timeNow.Year(), 12, 31, 23, 59, 59, 00, time.Local)
			break
		}
	default:
		{
			fmt.Println("============ INVALID")
			return startDate, endDate, errors.New("invalid time period")
		}
	}
	return startDate, endDate, nil
}
