package http

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/services"
)

type Handler struct {
	ginApp      *gin.Engine
	userService services.UserService
	TaskService services.TaskService
}

type HandlerOption func(*Handler)

func NewHandler(options ...HandlerOption) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithUserService(s services.UserService) HandlerOption {
	return func(handler *Handler) {
		handler.userService = s
	}
}

func WithTaskService(s services.TaskService) HandlerOption {
	return func(handler *Handler) {
		handler.TaskService = s
	}
}
func WithGinEngine(r *gin.Engine) HandlerOption {
	return func(handler *Handler) {
		handler.ginApp = r
	}
}

func (h *Handler) InitRoute() {
	h.ginApp.POST("/v1/user/login", h.Login)
	h.ginApp.GET("/v1/tasks", h.Authorise, h.ListTasks)
	h.ginApp.POST("/v1/tasks", h.Authorise, h.CreateTasks)
}
