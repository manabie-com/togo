package define

import "errors"

var (
	Unauthenticated = errors.New("request unauthenticated")
	FailedValidation = errors.New("validate request failed")
	DatabaseError = errors.New("error from database")
	CacheError = errors.New("error from cache")
	AccountNotAuthorized = errors.New("account not authorized")
	AccountNotExist = errors.New("account not exist")
	UserOverLimitTask = errors.New("user over limit task per day")
	Unknown = errors.New("unknown")
)
