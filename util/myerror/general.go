package myerror

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func ErrGRPC(err error) MyError {
	st := status.Convert(err)

	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: st.Code().String(),
		Message:   st.Message(),
	}
}

func ErrUnauthorized() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "000001",
		Message:   "Unauthorized.",
	}
}

func ErrMaxTodo() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusExpectationFailed,
		ErrorCode: "000002",
		Message:   "Too many requests.",
	}
}

func ErrInvalidParams(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "000003",
		Message:   "Invalid params.",
	}
}

func ErrGetUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "000003",
		Message:   "Failed to get user.",
	}
}
