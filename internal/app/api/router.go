package api

import (
	"net/http"
	"time"

	"github.com/dinhquockhanh/togo/internal/pkg/http/middleware"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	router := gin.New()
	router.Use(middleware.SetLogger(log.Root()))
	router.GET("/ping", ping)

	return router
}

func ping(context *gin.Context) {
	log.FromCtx(context.Request.Context()).Info("ping handler")
	time.Sleep(time.Second * 5)
	context.String(http.StatusOK, "pong")
}
