package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/transport/http/task_handlers"
	"net/http"
)

func ConfigTaskRouter(group *gin.RouterGroup, configs *config.ThirdAppAdapter) {
	NewRouterGroup(group, configs).
		Load(
			Route{
				JwtValidation: true,
				HandlerFunc:   task_handlers.HttpCreateTask,
				MethodName:    http.MethodPost,
				RelativePath:  "/tasks/",
			},
			Route{
				JwtValidation: true,
				HandlerFunc:   task_handlers.HttpListTasks,
				MethodName:    http.MethodGet,
				RelativePath:  "/tasks/",
			},
		)
}
