package errs

import "errors"

var (
	ErrInternalAppFailure = errors.New("app_failure")
	ErrDatabaseFailure    = errors.New("database_failure")
	ErrRecordNotFound     = errors.New("record_not_found")
)
