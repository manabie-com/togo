package http

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func HandleError(c *gin.Context, err error, data ...interface{}) bool {
	c.Header("Content-Type", "application/json")
	if err != nil {
		logInternalServerError(err, data)
		ErrBadRequest(c, err.Error())
		return true
	}
	return false
}

func logInternalServerError(err error, data interface{}) {
	logrus.WithFields(logrus.Fields{
		"data": data,
	}).Errorf("%+v", err)
}

func Success(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusOK, ResponseData{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data:    data,
	})
}

func ErrBadRequest(c *gin.Context, message ...string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Status:  http.StatusBadRequest,
		Message: strings.Join(message, ", "),
	})
}

func ErrNotFound(c *gin.Context, message ...string) {
	c.JSON(http.StatusNotFound, ResponseData{
		Status:  http.StatusNotFound,
		Message: strings.Join(message, ", "),
	})
}
