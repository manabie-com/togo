package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	yyyymmddRegex = `^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`
)

// validateNotEmpty check if the trimmed space of a string is not empty
func validateNotEmpty(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

// validateYYYYMMDD checks if the string is in format of YYYY-MM-DD
func validateYYYYMMDD(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(yyyymmddRegex, fl.Field().String())

	return matched
}
