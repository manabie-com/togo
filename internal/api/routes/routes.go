package routes

import (
	"example.com/m/v2/internal/api/handlers"
	"example.com/m/v2/internal/api/handlers/common"
	"example.com/m/v2/internal/api/handlers/task"
	"example.com/m/v2/internal/api/handlers/user"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine, service handlers.MainUseCase) {
	common.NewHandler(router, service)
	user.NewHandler(router, service)
	task.NewHandler(router, service)
}
