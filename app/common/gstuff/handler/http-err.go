package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// EchoError is format err for end user, return echo JSON
func EchoError(err error, c echo.Context) {

	// Tạo default response
	response := &ResponseContent{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}

	// Kiểm tra err is formatted with HTTPError or not
	if httpError, ok := err.(*HTTPError); ok {

		response.Code = httpError.HTTPstatus
		response.Message = httpError.ErrMsg
		response.CodeMessage = httpError.CodeMsg
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(response.Code)
		} else {
			c.JSON(response.Code, response)
		}
	}
}

// HTTPError is stored HTTP template response to client
type HTTPError struct {
	HTTPstatus int    `json:"http_status"`
	ErrMsg     string `json:"err_msg"`  // Message err description
	CodeMsg    string `json:"code_msg"` // Code err
	Language   string `json:"language"`
}

// NewHTTPError return a HTTPError
func NewHTTPError(httpStatus int, errMsg, codeMsg string, language ...string) *HTTPError {

	httpError := &HTTPError{
		HTTPstatus: httpStatus,
		ErrMsg:     errMsg,
		CodeMsg:    codeMsg,
		Language:   "vi",
	}
	// Cập nhật language
	if len(language) > 0 {
		httpError.Language = language[0]
	}

	return httpError
}

// Error : return err description message
func (h HTTPError) Error() string {
	dataByte, _ := json.Marshal(h)
	return string(dataByte)
}
