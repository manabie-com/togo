package domain

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

var (
	ErrInvalidArg = Error{
		Code:    "invalid_arg",
		Message: "invalid argument",
	}
	ErrNotFound = Error{
		Code:    "not_found",
		Message: "not found",
	}
)
