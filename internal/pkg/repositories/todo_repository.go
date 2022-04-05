package repositories

import (
	"context"
	"time"
	"togo/internal/pkg/domain/entities"
	"togo/pkg/utils"

	"gorm.io/gorm"
)

// TodoRepository interface
type TodoRepository interface {
	Create(ctx context.Context, todo entities.Todo) error
	CountTodosByDay(ctx context.Context, userID int) (int, error)
}

type todoRepository struct {
	DB *gorm.DB
}

// Create func
func (r *todoRepository) Create(ctx context.Context, todo entities.Todo) error {
	result := r.DB.WithContext(ctx).Create(&todo)
	return result.Error
}

// CountTodosByDay func
func (r *todoRepository) CountTodosByDay(ctx context.Context, userID int) (int, error) {
	now := time.Now()
	startOfDay := utils.StartOfDay(now)
	endOfDay := utils.EndOfDay(now)
	todos := []entities.Todo{}
	result := r.DB.WithContext(ctx).Where("user_id = ? and created_at > ? and created_at < ?", userID, startOfDay, endOfDay).Find(&todos)
	return len(todos), result.Error
}

func NewToDoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		DB: db,
	}
}
