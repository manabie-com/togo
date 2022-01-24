//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/internal/api"
	v1 "github.com/trinhdaiphuc/togo/internal/api/v1"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
	"github.com/trinhdaiphuc/togo/internal/repository"
	"github.com/trinhdaiphuc/togo/internal/service"
	taskservice "github.com/trinhdaiphuc/togo/internal/service/task"
	userservice "github.com/trinhdaiphuc/togo/internal/service/user"
)

var (
	infrastructureSet = wire.NewSet(infrastructure.NewDB)
	repositorySet     = wire.NewSet(repository.NewTaskRepository, repository.NewUserRepository)
	serviceSet       = wire.NewSet(service.NewService, userservice.NewUserService, taskservice.NewTaskService)
	handlerSet       = wire.NewSet(v1.NewTaskHandler, v1.NewUserHandler)
)

func InitializeServer() (*api.Server, func(), error) {
	panic(wire.Build(api.NewServer, configs.NewConfig, serviceSet, handlerSet, repositorySet, infrastructureSet))
}
