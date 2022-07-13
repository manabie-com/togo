package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(result interface{}) gin.H {
	return gin.H{
		"success": false,
		"error":   result,
	}
}
