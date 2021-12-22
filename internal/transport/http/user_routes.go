package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/transport/http/user_handlers"
	"net/http"
)

func ConfigUserRouter(group *gin.RouterGroup, configs *config.ThirdAppAdapter) {
	NewRouterGroup(group, configs).
		Load(
			Route{
				JwtValidation: false,
				HandlerFunc:   user_handlers.HttpLogin,
				MethodName:    http.MethodPost,
				RelativePath:  "/login/",
			},
			Route{
				JwtValidation: false,
				HandlerFunc:   user_handlers.HttpSignup,
				MethodName:    http.MethodPost,
				RelativePath:  "/signup/",
			},
		)
}
