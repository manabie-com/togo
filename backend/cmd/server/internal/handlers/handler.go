package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/httpserver/middleware"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/task/tasktransport/taskgin"
	"togo/modules/user/userstorage"
	"togo/modules/user/usertransport/ginuser"
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

		users := engine.Group("/api/users")
		{
			users.POST("/register", ginuser.Register(sc))
			users.POST("/login", ginuser.Login(sc))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := userstorage.NewSQLStore(db)
		middlewareAuth := middleware2.RequireAuth(store, sc)

		engine.PATCH("/api/users/setting", middlewareAuth, ginuser.UpdateLimit(sc))

		todos := engine.Group("/api/tasks", middlewareAuth)
		{
			todos.PATCH("/:id", taskgin.UpdateTask(sc))
			todos.DELETE("/:id", taskgin.DeleteTask(sc))
			todos.POST("", taskgin.CreateTask(sc))
			todos.GET("", taskgin.ListTasks(sc))
		}
	}
}
