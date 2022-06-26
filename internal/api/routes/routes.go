package routes

import (
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/api/handlers/common"
	"github.com/manabie-com/togo/internal/api/handlers/task"
	"github.com/manabie-com/togo/internal/api/handlers/user"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine, service handlers.MainUseCase) {
	common.NewHandler(router, service)
	user.NewHandler(router, service)
	task.NewHandler(router, service)
}
