package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/akaru/internal/handler"
)

func InitializeAPI() *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)

	return r
}
