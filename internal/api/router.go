package api

import (
	"github.com/gin-gonic/gin"
)

func InitializeAPI() *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the task endpoint",
		})
	})

	return r
}
