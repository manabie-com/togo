package server

import (
	"fmt"
	"net/http"
)

// HTTPError represents an error that occurred while handling a request
type HTTPError struct {
	Code     int    `json:"code"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

// NewHTTPError creates a new HTTPError instance
func NewHTTPError(code int, etype string, message ...string) *HTTPError {
	he := &HTTPError{Code: code, Type: etype}
	if len(message) > 0 {
		he.Message = message[0]
	} else {
		he.Message = http.StatusText(code)
	}
	return he
}

// Error generates error message and makes it compatible to error type
func (he *HTTPError) Error() string {
	if he.Internal == nil {
		return fmt.Sprintf("code=%d, type=%s, message=%s", he.Code, he.Type, he.Message)
	}
	return fmt.Sprintf("code=%d, type=%s, message=%s, internal=%v", he.Code, he.Type, he.Message, he.Internal)
}

// SetInternal sets actual internal error for more details
func (he *HTTPError) SetInternal(err error) *HTTPError {
	he.Internal = err
	return he
}
