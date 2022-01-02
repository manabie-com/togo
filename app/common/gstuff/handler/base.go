package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/app/common/config"

	validator "gopkg.in/go-playground/validator.v9"
)

var cfg = config.GetConfig()

// NewValidator ..
func NewValidator() *MyValidator {
	return &MyValidator{validator: validator.New()}
}

// MyValidator ..
type MyValidator struct {
	validator *validator.Validate
}

// Validate ..
func (cv *MyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// GetRequestID ..
func GetRequestID(c echo.Context) (requestID string) {
	if c.Get("reqID") != nil {
		requestID = c.Get("reqID").(string)
	}
	return
}

// GetHTTPCtx Tạo ctx từ echo context, có request ID
func GetHTTPCtx(c echo.Context) context.Context {
	var requestID string
	if c.Get("reqID") != nil {
		requestID = c.Get("reqID").(string)
	}
	httpCtx := context.WithValue(c.Request().Context(), "reqID", requestID)
	return httpCtx
}

// Health func
func Health(c echo.Context) (err error) {
	return c.JSON(Success(nil))
}

// ResponseContent struct
type ResponseContent struct {
	Code             int         `json:"code"`
	Message          string      `json:"message"`
	Data             interface{} `json:"data"`
	CodeMessage      string      `json:"code_message,omitempty"`
	CodeMessageValue string      `json:"code_message_value,omitempty"`
}

// Success func
func Success(data interface{}) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

// SuccessCustom func
func SuccessCustom(data interface{}, msg string) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:             http.StatusOK,
		Message:          "Success",
		Data:             data,
		CodeMessageValue: msg,
	}
}

// NotFound func
func NotFound(message string) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// ContextCanceled is check context is canceled or not
func ContextCanceled(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return fmt.Errorf("Request is canceled")
	default:
		return nil
	}
}

// ContextDeadline  is check context is deadline or not
func ContextDeadline(ctx context.Context) error {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return fmt.Errorf("Request is deadline")
	default:
		return nil
	}
}
