package util

import (
	"github.com/go-playground/validator"
)

// CustomValidator struct
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate func
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
