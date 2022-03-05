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

func ErrInvalidParams(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "000002",
		Message:   "Invalid params.",
	}
}

func ErrGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000003",
		Message:   "Failed to get.",
	}
}

func ErrCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000004",
		Message:   "Failed to create.",
	}
}

func ErrUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000005",
		Message:   "Failed to update.",
	}
}

func ErrDelete(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000006",
		Message:   "Failed to delete.",
	}
}

func ErrNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "000007",
		Message:   "Not found.",
	}
}
