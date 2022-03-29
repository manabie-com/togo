package routes

import (
	"togo/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/to-do")
	{
		api.GET("", controllers.Index)

		api.POST("", controllers.CreateTodo)
	}

	return r
}