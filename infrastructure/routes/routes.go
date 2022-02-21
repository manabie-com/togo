package routes

import (
	"github.com/gin-gonic/gin"
	"togo/interface/controllers"
)

func NewRouter(router *gin.Engine, ctl controllers.AppController) *gin.Engine {

	router.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"service": "API Togo",
			"status":  1,
		})
	})

	todoGroup := router.Group("api/v1/todo")
	{
		todoGroup.GET("/user", ctl.GetAllTodoUser)
	}

	return router
}
