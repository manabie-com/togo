package model

import "fmt"

type ErrorCode string

const (
	ErrUnknown          ErrorCode = "err.unknown"
	ErrInvalidUserModel ErrorCode = "err.model.user.invalid"

	ErrGetUser          ErrorCode = "err.store.users.get"
	ErrCreateUser       ErrorCode = "err.store.users.create"
	ErrAuthenticateUser ErrorCode = "err.model.user.authenticate"
	ErrListTasks        ErrorCode = "err.store.tasks.list"
	ErrAddTasks         ErrorCode = "err.store.tasks.add"
	ErrCountTasks       ErrorCode = "err.store.tasks.count"

	ErrTasksLimitExceeded ErrorCode = "err.usecase.users.task_limit_exceeded"
)

type Error struct {
	Code          ErrorCode `json:"code"`
	DetailedError string    `json:"-"` // Internal error string to help the developer
}

func NewError(code ErrorCode, detailedError string) *Error {
	return &Error{
		Code:          code,
		DetailedError: detailedError,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.DetailedError)
}

func GetErrorCode(err error) ErrorCode {
	if err == nil {
		return ErrUnknown
	}

	e, ok := err.(*Error)
	if !ok {
		return ErrUnknown
	}

	return e.Code
}
