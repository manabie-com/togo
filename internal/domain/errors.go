package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrTooMany             = errors.New("Exceed number of task today")
	ErrInvalidCredential   = errors.New("incorrect user_id/pwd")
)
