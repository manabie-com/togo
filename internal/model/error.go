package model

import "fmt"

type ErrorCode string
const (
	ErrInvalidUserModel             ErrorCode = "err.model.user.invalid"

	ErrGetUser ErrorCode = "err.store.user.get"
	ErrCreateUser ErrorCode = "err.store.user.create"
	ErrListTasks ErrorCode = "err.store.tasks.list"
	ErrAddTasks ErrorCode = "err.store.tasks.add"
)

type Error struct {
	Code             ErrorCode              `json:"code"`
	DetailedError    string                 `json:"-"` // Internal error string to help the developer
}

func NewError(code ErrorCode, detailedError string) *Error {
	return & Error{
		Code:          code,
		DetailedError: detailedError,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.DetailedError)
}
