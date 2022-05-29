package errdef

import "errors"

var (
	InvalidUsernameOrPassword = errors.New("Invalid username or password")
	TokenWrongFormat          = errors.New("Token is incorrect")

	DupplicateTask   = errors.New("Dupplicate task")
	LimitTaskCreated = errors.New("Limit task created in day")
	SystemError      = errors.New("System error")
)
