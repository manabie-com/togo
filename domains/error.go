package domains

import "errors"

var (
	ErrorNotFound = errors.New("not found")
	ErrorUnAuthorized = errors.New("un-authorized")
)
