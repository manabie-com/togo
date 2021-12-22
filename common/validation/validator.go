package validation

import (
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) map[string]interface{} {
	validate := validator.New()
	errors := map[string]interface{}{}

	if err := validate.Struct(data); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errors["base"] = err
			return errors
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
	}
	return errors
}
