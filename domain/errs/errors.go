package errs

import "errors"

var (
	ErrMaxToDoExceeded = errors.New("maximum todo exceeded")
)
