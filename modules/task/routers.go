package task

import (
	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/middleware/jwt"
	"github.com/khoale193/togo/modules/task/handler"
)

func TaskRouter(router *gin.RouterGroup) {
	router.Use(jwt.JWT())
	{
		router.POST("/task", handler.CreateTask)
	}
}
