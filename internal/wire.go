//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"togo/configs"
	"togo/internal/api"
	v1 "togo/internal/api/v1"
	"togo/internal/infrastructure"
	"togo/internal/repository"
	"togo/internal/service"
	taskservice "togo/internal/service/task"
	userservice "togo/internal/service/user"
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
