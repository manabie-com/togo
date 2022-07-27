package usecase

import "manabieAssignment/internal/core/entity"

type TodoUseCase interface {
	CreateTodo(todo entity.Todo) (uint, error)
}
