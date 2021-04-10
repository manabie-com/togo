package infra

import (
	"github.com/looplab/eventhorizon"
	"github.com/manabie-com/togo/internal/services/auth"
	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"
	"github.com/manabie-com/togo/internal/services/users"
)

func ProvideAuthService(cfg *AppConfig, userRepo users.UserRepo) auth.Service {
	return auth.NewAuthService(userRepo, cfg.SecretJWT)
}

func ProvideUserTaskService(
	commandBus eventhorizon.CommandHandler,
	userConfigRepo users.UserConfigRepo,
) user_tasks.Service {
	return user_tasks.NewService(commandBus, userConfigRepo)
}
