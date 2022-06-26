package common

import (
	"example.com/m/v2/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, service handlers.MainUseCase) {
	apiUser := router.Group("/")
	{
		apiUser.POST("/login", Login(service))
	}
}
