package handler

import (
	"github.com/gin-gonic/gin"
	"togo/middleware"
	taskrest "togo/module/task/transport/rest"
	userrest "togo/module/user/transport/rest"
	"togo/module/userconfig/transport/rest"
	"togo/server"
)

func RestHandler(sc server.ServerContext) func() *gin.Engine {
	return func() *gin.Engine {
		router := gin.Default()
		router.Use(middleware.Recovery())
		users := router.Group("/users")
		{
			users.POST("/tasks/add", taskrest.CreateTasks(sc))
			users.POST("", userrest.CreateUser(sc))
			users.PUT("/:id/tasks/update", rest.UpdateUserConfig(sc))
		}

		return router
	}
}
