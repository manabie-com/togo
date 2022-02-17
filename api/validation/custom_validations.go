package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	yyyymmddRegex = `^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`
)

func validateNotEmpty(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func validateYYYYMMDD(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(yyyymmddRegex, fl.Field().String())

	return matched
}
