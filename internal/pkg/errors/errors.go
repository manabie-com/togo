package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Error defines a standard application error.
type Error struct {
	Code    int    `json:"code"`             // Machine-readable error code
	Message string `json:"message"`          // Human-readable message
	Op      string `json:"-"`                // Logical operation (just log)
	Err     error  `json:"-"`                // Embedded error (just log)
	Detail  any    `json:"errors,omitempty"` // JSON encoded data
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Op, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func MakeValidationError(ve validator.ValidationErrors) *Error {
	errs := make([]ValidateError, len(ve))
	for i, fe := range ve {
		errs[i] = ValidateError{Field: fe.Field(), Message: getErrorMsg(fe)}
	}
	return &Error{
		Code:    http.StatusBadRequest,
		Message: "validate error",
		Detail:  errs,
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is missing"
	case "email":
		return "invalid email"
	case "oneof":
		return fmt.Sprintf("the value must be one of [%v]", strings.Replace(fe.Param(), " ", ", ", -1))
	case "min":
		return fmt.Sprintf("the value must be greater than or equal with %v", fe.Param())
	case "max":
		return fmt.Sprintf("the value must be less than or equal with %v", fe.Param())
	}
	// TODO: handle another tag

	return "invalid"
}

func NewNotFoundErr(err error, op string, domain string) *Error {
	return &Error{Op: op, Err: err, Code: http.StatusNotFound, Message: fmt.Sprintf("Not found %s", domain)}
}
