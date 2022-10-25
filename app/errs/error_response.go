package errs

import "github.com/ansidev/togo/gingo/validation"

type ErrorResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Error   string                `json:"error,omitempty"`
	Errors  []validation.ErrorMsg `json:"errors,omitempty"`
}
