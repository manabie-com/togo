package cmerrors

import (
	"errors"
)

var (
	ErrUserAlreadyExist = errors.New("User already exists")
	ErrUserNotFound     = errors.New("User not found")
	ErrPasswordNotMatch = errors.New("Password not match")

	ErrTooManyTask  = errors.New("Too many tasks")
	ErrTaskNotFound = errors.New("Task Not found")
)
