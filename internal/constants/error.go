package constants

import "errors"

var (
	ErrIncorrectUserIdOrPassword = errors.New("incorrect user_id/pwd")
	ErrCreateToken               = errors.New("create token failed")
	ErrGetUserFromContext        = errors.New("get user from context failed")
	ErrMaximumCreatedTask        = errors.New("maximum created task")
)

var (
	KeyUserId        = "user-id"
	KeyAuthorization = "Authorization"
)
