package define

import "errors"

var (
	Unauthenticated = errors.New("request unauthenticated")
	FailedValidation = errors.New("validate request failed")
	DatabaseError = errors.New("error from database")
	AccountNotAuthorized = errors.New("account not authorized")
	AccountNotExist = errors.New("account not exist")
	Unknown = errors.New("unknown")
)
