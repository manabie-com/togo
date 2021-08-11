package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/manabie-com/togo/common/database"
	"github.com/manabie-com/togo/common/middleware"
	"github.com/manabie-com/togo/domain"
	"github.com/manabie-com/togo/handler"
	"github.com/manabie-com/togo/repo"
	"net/http"
)

func Serve() {
	appCfg := &AppConfig{}
	err := configor.Load(appCfg, "app.yml")
	if err != nil {
		panic(err)
	}
	db := database.NewDB(appCfg.Db)

	userRepo := repo.NewUserRepo()
	taskRepo := repo.NewTaskRepo()
	todoService := domain.NewTodoService(appCfg.Token.Key, appCfg.Token.Timeout, userRepo, taskRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	route := gin.Default()
	route.Use(middleware.Transaction(db.GetDb()))
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"heal_check": "ok",
		})
	})

	route.GET("/login", todoHandler.GetAccessToken)

	routeAccess := route.Group("/tasks")
	routeAccess.Use(middleware.AccessToken(appCfg.Token.Key))
	routeAccess.POST("", todoHandler.CreateTask)
	routeAccess.GET("", todoHandler.GetTasks)

	route.Run(fmt.Sprintf("%v:%v", appCfg.Host.Host, appCfg.Host.Port))
}
