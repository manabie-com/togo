package tasks

import (
	"github.com/manabie-com/togo/cmd/middlewares"
	"github.com/manabie-com/togo/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiTask := router.Group("/")
	apiTask.Use(middlewares.ValidateToken())
	{
		apiTask.POST("/tasks", AddTask(service))
	}
}
