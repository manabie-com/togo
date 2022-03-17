package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NewValidator returns the *validator.Validate type with preset configurations
func NewValidator() *validator.Validate {
	v := validator.New()

	// Set the field name of the error as the json tag of the struct
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom validations
	v.RegisterValidation("not_empty", validateNotEmpty)
	v.RegisterValidation("yyyymmdd_date", validateYYYYMMDD)

	return v
}
