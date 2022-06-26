package task

import (
	"github.com/manabie-com/togo/cmd/middlewares"
	"github.com/manabie-com/togo/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiTask := router.Group("/tasks")
	apiTask.Use(middlewares.ValidateToken())
	{
		apiTask.POST("/", AddTask(service))
	}
}
