package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(result interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    result,
	}
}
