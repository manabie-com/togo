package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/httpserver/middleware"
	"github.com/phathdt/libs/togo_appgrpc"
	"togo/common"
	"togo/modules/task/tasktransport/taskgin"
	middleware2 "togo/plugin/middleware"
)

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func NewHandlers(sc goservice.ServiceContext) func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.Use(middleware.Recover())

		engine.GET("/ping", ping())

		userService := sc.MustGet(common.PluginGrpcUserClient).(togo_appgrpc.UserClient)
		middlewareAuth := middleware2.RequireAuth(userService, sc)

		todos := engine.Group("/api/tasks", middlewareAuth)
		{
			todos.PATCH("/:id", taskgin.UpdateTask(sc))
			todos.DELETE("/:id", taskgin.DeleteTask(sc))
			todos.POST("", taskgin.CreateTask(sc))
			todos.GET("", taskgin.ListTasks(sc))
		}
	}
}
