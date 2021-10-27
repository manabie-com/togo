package handler

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/internal/core/port"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	Begin(ctx context.Context, address string) error
}

func NewHttpHandler(taskService port.TaskService, jwtService port.JwtService) HttpHandler {
	return &httpHandler{
		taskService: taskService,
		jwtService:  jwtService,
	}
}

type httpHandler struct {
	taskService port.TaskService
	jwtService  port.JwtService
}

func (p *httpHandler) Begin(ctx context.Context, address string) error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}), p.ignoreOptionsMethod())
	p.setup(router.Group(""))

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (p *httpHandler) ignoreOptionsMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func (p *httpHandler) setup(router *gin.RouterGroup) {
	router.POST("/login", p.login)
	router.GET("/tasks", p.getTasks)
	router.POST("/tasks", p.addTask)
}
