package repositories

import (
	"context"
	"fmt"
	"togo/database/datastore"
	"togo/domain/models"
)

type TodoRepository struct{}

var todoRepository *TodoRepository

func NewTodoRepository() *TodoRepository {
	if todoRepository == nil {
		todoRepository = &TodoRepository{}
	}
	return todoRepository
}

func (tr *TodoRepository) CreateTodo(ctx context.Context, newTodo *models.Todo) error {
	if err := datastore.GetDB().WithContext(ctx).Create(newTodo).Error; err != nil {
		return fmt.Errorf("error TodoRepository.CreateTodo %v", err)
	}
	return nil
}
