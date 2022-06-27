package usecase

import (
	"errors"
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

func (t *todoUseCase) CreateTodo(todo entity.Todo) error {
	if err := todo.Validate(); err != nil {
		return err
	}
	numOfTodos, err := t.userRepo.CountTodosByDay(todo.UserID, todo.CreatedAt)
	if err != nil {
		return err
	}
	if numOfTodos > 5 {
		return errors.New("too many todos")
	}
	err = t.todoRepo.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}
