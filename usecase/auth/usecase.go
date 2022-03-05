package auth

import (
	"github.com/khangjig/togo/config"
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/repository/user"
)

type UseCase struct {
	UserRepo user.Repository
	Config   *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
		Config:   config.GetConfig(),
	}
}
