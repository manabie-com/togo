package errors

import "fmt"

var (
	ErrUserIdIsInvalid  = fmt.Errorf("user id is invalid")
	ErrUserDoesNotExist = fmt.Errorf("user does not exist")
	ErruserAlreadyExist = fmt.Errorf("user already exists")

	ErrTaskLimitExceeded = fmt.Errorf("task limit exceeded")
)
