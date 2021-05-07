package user

import (
	"github.com/gin-gonic/gin"
	jwt_service "manabie-com/togo/config/jwt-service"
	"manabie-com/togo/global"
)

func RegisterRoutes(route *gin.Engine) {
	userGroup := route.Group(global.Config.Prefix + "/v1")
	{
		userGroup.POST("/users/log_in", LogIn)
		userGroup.POST("/users", CreateOne)
		userGroup.GET("/users/tasks", jwt_service.DoAuthenticate(), ListTasks)
		userGroup.POST("/users/tasks", jwt_service.DoAuthenticate(), AddTask)
	}
}
