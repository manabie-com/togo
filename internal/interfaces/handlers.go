package interfaces

import (
	"database/sql"
	"path"

	"github.com/datshiro/togo-manabie/internal/interfaces/consts"
	"github.com/datshiro/togo-manabie/internal/interfaces/service/cache"
	task_handler "github.com/datshiro/togo-manabie/internal/interfaces/task"
	task_usecase "github.com/datshiro/togo-manabie/internal/interfaces/task/usecase"
	user_handler "github.com/datshiro/togo-manabie/internal/interfaces/user"
	user_usecase "github.com/datshiro/togo-manabie/internal/interfaces/user/usecase"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string, dbc *sql.DB, cacheS cache.CacheService) {
	taskUC := task_usecase.NewTaskUseCase(dbc, cacheS)
	userUC := user_usecase.NewUserUseCase(dbc)
	task_handler.RegisterHandlers(e, path.Join(apiPrefix, consts.TaskPath), taskUC, userUC)
	user_handler.RegisterHandlers(e, path.Join(apiPrefix, consts.UserPath), userUC)
}
