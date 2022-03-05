package todo

import (
	"github.com/khangjig/togo/config"
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/repository/todo"
	"github.com/khangjig/togo/repository/user"
)

type UseCase struct {
	UserRepo      user.Repository
	TodoRepo      todo.Repository
	TodoCacheRepo todo.CacheRepository
	Config        *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo:      repo.User,
		TodoRepo:      repo.Todo,
		TodoCacheRepo: repo.TodoCache,
		Config:        config.GetConfig(),
	}
}
