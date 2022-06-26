package user

import (
	"example.com/m/v2/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiUser := router.Group("/users")
	{
		apiUser.POST("/", CreateUser(service))
	}
}
