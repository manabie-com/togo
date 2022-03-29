package validator

import (
	"github.com/go-playground/validator/v10"
)


var v *validator.Validate

type ErrorJSON struct {
	Namespace string `json:"field"`
	ActualTag string `json:"validate"`
}

func Validate(i interface{}) []ErrorJSON {
	if v == nil {
		v = validator.New()
	}
	err := v.Struct(i)
	if (err != nil) {
		errJSONS := []ErrorJSON{}
		for _, err := range err.(validator.ValidationErrors) {
			errJSONS = append(errJSONS, ErrorJSON{ Namespace: err.Namespace(), ActualTag: err.ActualTag()})
		}

		return errJSONS
	}

	return nil
}