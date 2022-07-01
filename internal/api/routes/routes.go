package routes

import (
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/api/handlers/common"
	"github.com/manabie-com/togo/internal/api/handlers/tasks"
	"github.com/manabie-com/togo/internal/api/handlers/users"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine, service handlers.MainUseCase) {
	common.NewHandler(router, service)
	users.NewHandler(router, service)
	tasks.NewHandler(router, service)
}
