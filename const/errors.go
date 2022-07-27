package _const

import "errors"

var (
	ErrOnCreateUser = errors.New("Create user fail")
	ErrOnCreateTask = errors.New("Create task fail")
	ErrLimitedTask  = errors.New("Reach limit task")
	ErrUserNotFound = errors.New("User not found")
	ErrDateType     = errors.New("Date type not right")
	ErrValidate     = errors.New("Validate fail")
)
