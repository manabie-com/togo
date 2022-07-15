package errors

import (
	"fmt"
)

type SystemError int

const (
	InvalidRequest    SystemError = 400
	InternalServerErr SystemError = 500
)

var SystemErrorMap = map[SystemError]string{
	InvalidRequest:    "Invalid request",
	InternalServerErr: "Internal server error",
}

func (e SystemError) Error() string {
	return SystemErrorMap[e]
}

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

func NewParamErr(err string) CustomParamError {
	return CustomParamError(err)
}

type CustomParamError string

func (e CustomParamError) Error() string {
	return fmt.Sprintf("%s: %s", InvalidParamError, string(e))
}
