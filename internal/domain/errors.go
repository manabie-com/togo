package domain

import "errors"

var (
	ErrorMaximumTaskPerDay = errors.New("create limit task this day")
	UserNotFound           = errors.New("user not found")
	WrongPassword          = errors.New("wrong password")
	Unauthorized           = errors.New("unauthorized")
)
