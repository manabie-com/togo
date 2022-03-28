package domain

import (
	"fmt"
)

type Error struct {
	Code    string
	Message string
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

// Is returns true when error is an instance of domain.Error
// and has the same Error.Code
func (err Error) Is(target error) bool {
	domainErr, ok := target.(Error)
	if !ok {
		return false
	}
	return err.Code == domainErr.Code
}

// defines error groups for easier error handling at the infra layer
var (
	ErrInvalidArg = Error{
		Code:    CodeInvalidArg,
		Message: "invalid argument",
	}
	ErrNotFound = Error{
		Code:    CodeNotFound,
		Message: "not found",
	}
)

func NewError(code, msg string) error {
	return Error{Code: code, Message: msg}
}

func NotFound(msg string) error   { return NewError(CodeNotFound, msg) }
func InvalidArg(msg string) error { return NewError(CodeInvalidArg, msg) }

const (
	CodeInternal   = "internal"
	CodeNotFound   = "not_found"
	CodeInvalidArg = "invalid_argument"
)
