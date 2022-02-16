package apierrors

import "errors"

var (
	UserAlreadyExists = errors.New("User already exists")
	UserDoesNotExists = errors.New("User does not exists")
	MaxTasksReached   = errors.New("Max tasks created today has been reached")
)
