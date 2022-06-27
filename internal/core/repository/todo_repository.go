package repository

import "manabieAssignment/internal/core/entity"

type TodoRepository interface {
	CreateTodo(todo entity.Todo) error
}
