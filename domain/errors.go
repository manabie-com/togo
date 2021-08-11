package domain

import "errors"

var (
	ErrFailPrecondition = errors.New("fail precondition")
	ErrUserNotFound     = errors.New("can't found user")
)
