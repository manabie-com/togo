package myerror

import "net/http"

func ErrTodoTitleInvalid(message string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "200001",
		Message:   message,
	}
}

func ErrTodoContentInvalid(message string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "200002",
		Message:   message,
	}
}

func ErrTodoStatusInvalid() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "200002",
		Message:   "Invalid status",
	}
}
