package router

import (

	"github.com/gin-gonic/gin"
	"togo/api"
	"togo/utils"
)

func initTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(utils.JWTAuth())
	{
		taskRouter.GET("", api.GetAllTask)
		taskRouter.POST("", api.AddTask)
		taskRouter.PUT("/:id", api.UpdateTask)
		taskRouter.DELETE("/:id", api.DeleteTask)
	}

}
