package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Empty struct {}

// Response : JSON Response Object
type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseJSON ...
func ResponseJSON(c *gin.Context, d interface{}) {
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    d,
	})
}

// ResponseError ...
func ResponseError(c *gin.Context, s int, e string) {
	c.JSON(s, &Response{
		Success: false,
		Error:   e,
	})
	c.Abort()
}
