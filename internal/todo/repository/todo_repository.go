package repository

import (
	"gorm.io/gorm"
	"manabieAssignment/internal/core/entity"
	"manabieAssignment/internal/core/repository"
	"manabieAssignment/internal/todo/repository/dao"
)

type todoRepository struct {
	gormDB *gorm.DB
}

func NewTodoRepository(gormDB *gorm.DB) repository.TodoRepository {
	return &todoRepository{
		gormDB: gormDB,
	}
}

func (t *todoRepository) CreateTodo(todo entity.Todo) (uint, error) {
	todoDao := dao.Todo{
		UserID:  todo.UserID,
		Name:    todo.Name,
		Content: todo.Content,
		Model: gorm.Model{
			CreatedAt: todo.CreatedAt,
		},
	}
	err := t.gormDB.Create(&todoDao).Error
	if err != nil {
		return 0, err
	}
	return todoDao.ID, nil
}
