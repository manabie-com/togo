package usecase

import (
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/usecase/auth"
	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/usecase/user"
)

type UseCase struct {
	Auth auth.IUseCase
	User user.IUseCase
	Todo todo.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Auth: auth.New(repo),
		User: user.New(repo),
		Todo: todo.New(repo),
	}
}
