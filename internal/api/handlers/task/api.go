package task

import (
	"example.com/m/v2/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiTask := router.Group("/tasks")
	{
		apiTask.POST("/", AddTask(service))
	}
}
