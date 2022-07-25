package utils

import "net/http"

type Error struct {
	StatusCode int
	Message    string
}

// Custom error
func (e *Error) Error() string {
	return e.Message
}

// NewError /**
func NewError(c int, s string) Error {
	return Error{c, s}
}

// PanicInternalServerError /**
func PanicInternalServerError(err ...error) {
	message := "Something went wrong."
	if len(err) > 0 {
		message = err[0].Error()
	}
	panic(NewError(http.StatusInternalServerError, message))
}

// PanicBadRequest /**
func PanicBadRequest(err ...error) {
	message := "Bad request."
	if len(err) > 0 {
		message = err[0].Error()
	}
	panic(NewError(http.StatusBadRequest, message))
}

// PanicTooManyRequests /**
func PanicTooManyRequests() {
	panic(NewError(http.StatusTooManyRequests, "Too many requests."))
}

// PanicNotFound /**
func PanicNotFound() {
	panic(NewError(http.StatusNotFound, "Not found."))
}

// PanicUnauthorized /**
func PanicUnauthorized() {
	panic(NewError(http.StatusUnauthorized, "Unauthorized."))
}
