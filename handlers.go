package main

import (
	"github.com/gin-gonic/gin"
	appctx "github.com/manabie-com/togo/app_ctx"
	"github.com/manabie-com/togo/auth/transport"
	"github.com/manabie-com/togo/middleware"
	transport2 "github.com/manabie-com/togo/task/transport"
)

func setupHandlers(engine *gin.Engine, appctx appctx.IAppCtx) {
	engine.Use(middleware.Recovery())

	v1 := engine.Group("/v1/api")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", transport.UserLogin(appctx))
	}

	task := v1.Group("/tasks", middleware.Authenticator(appctx))
	{
		task.POST("", transport2.CreateTask(appctx))
	}
}
