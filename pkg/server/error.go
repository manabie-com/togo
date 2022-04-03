package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	InternalErrorType   = "INTERNAL"
	GenericErrorType    = "GENERIC"
	ValidationErrorType = "VALIDATION"
)

var ValidationErrors = map[string]string{
	"required": " is required",
	"email":    "'s value should be a valid email address",
}

// ErrorResponse represents the error response
type ErrorResponse struct {
	Error *HTTPError `json:"error"`
}

// HTTPError represents an error that occurred while handling a request
type HTTPError struct {
	Code     int    `json:"code"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

// ErrorHandler represents the custom http error handler
type ErrorHandler struct {
	e *echo.Echo
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

// NewHTTPInternalError creates a new HTTPError instance for internal error
func NewHTTPInternalError(message string) *HTTPError {
	return &HTTPError{Code: http.StatusInternalServerError, Type: InternalErrorType, Message: message}
}

// NewHTTPValidationError creates a new HTTPError instance for validation error
func NewHTTPValidationError(message string) *HTTPError {
	return &HTTPError{Code: http.StatusBadRequest, Type: ValidationErrorType, Message: message}
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

// newErrorHandler returns the ErrorHandler instance
func newErrorHandler(e *echo.Echo) *ErrorHandler {
	return &ErrorHandler{e}
}

// HandlerFunc is a centralized HTTP error handler.
func (h *ErrorHandler) HandlerFunc(err error, c echo.Context) {
	httpErr := NewHTTPError(http.StatusInternalServerError, InternalErrorType)

	switch e := err.(type) {
	case *HTTPError:
		if e.Code != 0 {
			httpErr.Code = e.Code
		}
		if e.Type != "" {
			httpErr.Type = e.Type
		} else {
			httpErr.Type = GenericErrorType
		}
		if e.Message != "" {
			httpErr.Message = e.Message
		}
		if e.Internal != nil {
			httpErr.Internal = e.Internal
		}
	case *echo.HTTPError:
		httpErr.Code = e.Code
		httpErr.Type = GenericErrorType
		switch em := e.Message.(type) {
		case string:
			httpErr.Message = em
		case []string:
			httpErr.Message = strings.Join(em, "\n")
		case map[string]interface{}:
			if jsonStr, err := json.Marshal(em); err == nil {
				httpErr.Message = string(jsonStr)
			}
		default:
			httpErr.Message = fmt.Sprintf("%+v", em)
		}
		if e.Internal != nil {
			httpErr.Internal = e.Internal
		}
	case validator.ValidationErrors:
		httpErr.Code = http.StatusBadRequest
		httpErr.Type = ValidationErrorType
		var errMsg []string
		for _, v := range e {
			errMsg = append(errMsg, getVldErrorMsg(v))
		}
		httpErr.Message = strings.Join(errMsg, "\n")
	default:
		if h.e.Debug {
			httpErr.Message = err.Error()
		}
	}

	// Send error response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(httpErr.Code)
		} else {
			err = c.JSON(httpErr.Code, ErrorResponse{Error: httpErr})
		}
		if err != nil {
			h.e.Logger.Error(err)
		}
	}
}

func getVldErrorMsg(v validator.FieldError) string {
	field := v.Field()
	vtag := v.ActualTag()
	vtagVal := v.Param()

	if msg, ok := ValidationErrors[vtag]; ok {
		return field + msg
	}

	switch vtag {
	case "oneof":
		return field + " should be one of " + strings.Replace(vtagVal, " ", ", ", -1)
	case "ltfield":
		return field + " should be less than " + vtagVal
	case "gtfield":
		return field + " should be greater than " + vtagVal
	case "eqfield":
		return field + " does not match " + vtagVal
	}

	return field + " failed on " + vtag + " validation"
}
