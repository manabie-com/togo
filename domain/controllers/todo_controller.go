package controllers

import (
	"context"
	"togo/domain/models"
	"togo/domain/services"
)

type todoController struct {
	todoService services.TodoService
}

type TodoController interface {
	CreateTodo(ctx context.Context, newTodo *models.NewTodo) (*models.OverviewTodo, error)
}

func (tc *todoController) CreateTodo(ctx context.Context, newTodo *models.NewTodo) (*models.OverviewTodo, error) {
	todo, err := tc.todoService.CreateTodo(ctx, newTodo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func NewTodoController(todoService services.TodoService) *todoController {
	return &todoController{
		todoService: todoService,
	}
}
