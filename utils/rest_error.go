package utils

import "net/http"

type ErrorRest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GeneralError(code int, message string) *ErrorRest {
	return &ErrorRest{
		Code:    code,
		Message: message,
	}
}

func ErrBadRequest(message string) *ErrorRest {
	return &ErrorRest{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ErrInternal(message string) *ErrorRest {
	return &ErrorRest{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
