package apperror

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"-"`
	ErrorKey   string `json:"errorKey"`
	Message    string `json:"message"`
	RootErr    error  `json:"-"`
}

func NewAppError(root error, statusCode int, msg, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		ErrorKey:   key,
	}
}

func NewErrorResponse(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		ErrorKey:   key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}
