package xerrors

import "fmt"

type Code int

const (
	NoError         = Code(0)
	Unknown         = Code(1)
	InvalidArgument = Code(2)
	Internal        = Code(3)
	UnAuthorized    = Code(4)
)

type XError struct {
	Code    Code
	Message string
	Err     error
}

func (e XError) Error() string {
	return e.Message
}

func Error(code Code, err error) XError {
	xError := XError{
		Code: code,
		Err:  err,
	}
	if err != nil {
		xError.Message = err.Error()
	}
	return xError
}

func ErrorM(code Code, err error, msg string) XError {
	return XError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func ErrorMf(code Code, err error, msg string, args ...interface{}) XError {
	return XError{
		Code:    code,
		Message: fmt.Sprintf(msg, args),
		Err:     err,
	}
}
