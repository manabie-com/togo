package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	maxErrMsg      = "Length of field '%s' may not be greater than %v characters"
	requiredErrMsg = "Field '%s' is required"
	fieldErrMsg    = "Field validation for '%s' failed on the '%s' tag"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "max":
		return fmt.Sprintf(maxErrMsg, fe.Field(), fe.Param())
	case "required":
		return fmt.Sprintf(requiredErrMsg, fe.Field())
	default:
		return fmt.Sprintf(fieldErrMsg, fe.Field(), fe.Tag())
	}
}
