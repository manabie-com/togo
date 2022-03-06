package domain

import (
	"github.com/triet-truong/todo/todo/dto"
	"github.com/triet-truong/todo/todo/model"
)

type TodoUseCase interface {
	AddTodo(newAlert dto.TodoDto) error
}

type TodoRepository interface {
	InsertItem(item model.TodoItemModel) error
	SelectUser(id uint) (model.UserModel, error)
}

type TodoCacheRepository interface {
	SetUser(user model.UserRedisModel) error
	GetCachedUser(id uint) (model.UserRedisModel, error)
}
