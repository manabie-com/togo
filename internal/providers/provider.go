//+build wireinject

package providers

import (
	"github.com/google/wire"
	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/middlewares"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"gorm.io/gorm"
)

func ProvideApplicationMiddleware(jwtSecretKey string) *middlewares.AppMiddleware {
	wire.Build(
		middlewares.NewAppMiddleware,
		helpers.NewTokenProvider,
	)

	return &middlewares.AppMiddleware{}
}

func ProvideTaskHandler(db *gorm.DB) *handlers.TaskHandler {
	wire.Build(
		handlers.NewTaskHandler,
		services.NewTaskService,
		services.NewConfigurationService,
		repositories.NewTaskRepository,
		repositories.NewConfigurationRepository,
	)

	return &handlers.TaskHandler{}
}

func ProvideUserHandler(db *gorm.DB, jwtSecretKey string) *handlers.UserHandler {
	wire.Build(
		handlers.NewUserHandler,
		services.NewUserService,
		helpers.NewTokenProvider,
		repositories.NewUserRepository,
	)

	return &handlers.UserHandler{}
}

func ProvideConfigurationHandler(db *gorm.DB) *handlers.ConfigurationHandler {
	wire.Build(
		handlers.NewConfigurationHandler,
		services.NewConfigurationService,
		repositories.NewConfigurationRepository,
	)

	return &handlers.ConfigurationHandler{}
}
