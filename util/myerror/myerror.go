package myerror

import (
	"strings"

	"github.com/pkg/errors"
)

type MyError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
}

func (e MyError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

func (e MyError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string) MyError {
	return MyError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
	}
}
