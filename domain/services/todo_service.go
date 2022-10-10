package services

import (
	"context"
	"time"
	"togo/domain/models"
	"togo/domain/repositories"
)

type TodoService interface {
	CreateTodo(ctx context.Context, newTodo *models.NewTodo) (*models.OverviewTodo, error)
}

type todoService struct {
	todoRepository *repositories.TodoRepository
}

func (ts *todoService) CreateTodo(ctx context.Context, newTodo *models.NewTodo) (*models.OverviewTodo, error) {
	var todo *models.Todo
	todo.Text = newTodo.Text
	todo.Title = newTodo.Title
	todo.Done = newTodo.Done
	todo.FkUser = newTodo.FkUser

	err := ts.todoRepository.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &models.OverviewTodo{
		ID:        0,
		Title:     "",
		Text:      "",
		Done:      "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func NewTodoService(
	todoRepository *repositories.TodoRepository,
) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}
