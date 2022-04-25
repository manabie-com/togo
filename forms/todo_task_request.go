package forms

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type TodoTaskRequest struct {
	UserId     int    `json:"user_id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Detail     string `json:"detail"`
	RemindDate string `json:"remind_date"`
}

// Validate struct function
func (req TodoTaskRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, validatorErr := range err.(validator.ValidationErrors) {
			return fmt.Errorf("invalid request in field [%+v], tag [%+v], value [%+v]",
				validatorErr.StructNamespace(), validatorErr.Tag(), validatorErr.Param())
		}
	}
	return nil
}
