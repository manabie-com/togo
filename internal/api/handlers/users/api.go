package users

import (
	"github.com/manabie-com/togo/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiUser := router.Group("/users")
	{
		apiUser.POST("/", CreateUser(service))
	}
}
