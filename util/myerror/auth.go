package myerror

import "net/http"

func ErrInvalidEmailPassword() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "100001",
		Message:   "Invalid email/password.",
	}
}

func ErrEncodeToken(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "100002",
		Message:   "Failed to encode token.",
	}
}

func ErrRegexp(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "100003",
		Message:   err.Error(),
	}
}

func ErrGetTotalRequest(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "100004",
		Message:   err.Error(),
	}
}

func ErrSetTotalRequest(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "100005",
		Message:   err.Error(),
	}
}
