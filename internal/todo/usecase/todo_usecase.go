package usecase

import (
	"manabieAssignment/internal/core/entity"
	"manabieAssignment/internal/core/repository"
	"manabieAssignment/internal/core/usecase"
)

type todoUseCase struct {
	todoRepo repository.TodoRepository
	userRepo repository.UserRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository, userRepo repository.UserRepository) usecase.TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}

func (t *todoUseCase) CreateTodo(todo entity.Todo) (uint, error) {
	if err := todo.Validate(); err != nil {
		return 0, err
	}
	err := t.userRepo.IsUserExisted(todo.UserID)
	if err != nil {
		return 0, err
	}
	err = t.userRepo.IsUserHavingMaxTodo(todo.UserID, todo.CreatedAt)
	if err != nil {
		return 0, err
	}
	todoId, err := t.todoRepo.CreateTodo(todo)
	if err != nil {
		return 0, err
	}
	return todoId, err
}
