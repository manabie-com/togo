package helper

import (
	"regexp"
)

// DateValidator regex date validator
func DateValidator(date string) bool {
	r, _ := regexp.Compile(`\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])*`)
	return r.MatchString(date)
}
