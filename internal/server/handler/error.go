package handler

import "errors"

var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrEmptyRequest    = errors.New("request should not be empty")
	ErrEmptyTitle      = errors.New("title should not be empty")
	ErrTooManyRequests = errors.New("too many requests")
)
