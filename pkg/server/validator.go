package server

import "github.com/go-playground/validator/v10"

// Validator holds custom validator
type Validator struct {
	validator *validator.Validate
}

// Validate validates the request
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// newValidator creates new custom validator
func newValidator() *Validator {
	validator := validator.New()

	return &Validator{validator: validator}
}
