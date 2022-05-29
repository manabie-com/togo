package api

import (
	"net/http"
	"time"

	"github.com/dinhquockhanh/togo/internal/app/auth"
	"github.com/dinhquockhanh/togo/internal/pkg/http/middleware"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *Handler) (http.Handler, error) {
	router := gin.New()

	router.Use(middleware.Recover())
	router.Use(middleware.SetLogger(log.Root()))
	router.Use(auth.SetUserMiddleware(h.auth.Tokenizer()))
	router.GET("/ping", ping)

	v1 := router.Group("/api/v1")
	{
		v1Auth := v1.Group("/").Use(auth.RequireAuth())

		v1Auth.GET("/tasks/:id", h.task.GetByID)
		v1Auth.POST("/tasks", h.task.CreateTask)
		v1Auth.PATCH("/tasks", h.task.AssignTask)

		v1Auth.GET("/users/:username", h.user.GetByUserName)

		v1.POST("/users", h.user.CreateUser)
		v1.POST("/auth/login", h.auth.Login)

	}

	return router, nil
}

func ping(context *gin.Context) {
	log.FromCtx(context.Request.Context()).Info("ping handler")
	time.Sleep(time.Second * 5)
	context.String(http.StatusOK, "pong")
}
