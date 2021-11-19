package routers

import (
	"manabie/manabie/controllers"

	"github.com/gin-gonic/gin"
)

func TodoAPIRoute(r *gin.RouterGroup) {
	r.GET("/tasks", controllers.GetAllTask)
	r.POST("/tasks", controllers.CreateTask)
}
