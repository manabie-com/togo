package validate

import (
	"github.com/go-playground/validator/v10"
)

var v = newCustomValidate()

type customValidate struct {
	validate *validator.Validate
}

// newCustomValidate customValidate constructor
func newCustomValidate() *customValidate {
	v := validator.New()

	// any custom for validate

	return &customValidate{
		validate: v,
	}
}

// Struct validate struct
func Struct(request interface{}) error {
	return v.validate.Struct(request)
}
