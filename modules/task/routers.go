package task

import (
	"github.com/gin-gonic/gin"

	"togo/middleware/jwt"
	"togo/modules/task/handler"
)

func TaskRouter(router *gin.RouterGroup) {
	router.Use(jwt.JWT())
	{
		router.POST("/level", handler.CreateTask)
	}
}
