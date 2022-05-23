package domain

import "errors"

var (
	ErrUserNotFound    = errors.New("user id not found")
	ErrCreateTask      = errors.New("failed to create task")
	ErrExceedTaskLimit = errors.New("exceed task limit")
)
