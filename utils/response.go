package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	StatusCode int
	Message    string
	ErrorMsg   string
}

func ResponseJson(c *gin.Context, statusCode int, data interface{}) {
	c.IndentedJSON(statusCode, data)
}
