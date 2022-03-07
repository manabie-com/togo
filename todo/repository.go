package todo

import (
	"github.com/triet-truong/todo/todo/model"
)

type TodoRepository interface {
	InsertItem(item model.TodoItemModel) error
	SelectUser(id uint) (model.UserModel, error)
}

type TodoCacheRepository interface {
	SetUser(user model.UserRedisModel) error
	GetCachedUser(id uint) (model.UserRedisModel, error)
}
