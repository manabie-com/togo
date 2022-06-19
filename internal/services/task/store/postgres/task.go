package postgres

import (
	"togo/internal/pkg/db"
	"togo/internal/services/task/domain"
	"togo/internal/services/task/store"
)

type TaskRepository struct {
	DB *db.DB
}

func NewTaskRepository(db *db.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) Save(entity *domain.Task) error {
	_, err := r.DB.Conn.Model(entity).Returning("*").Insert()
	return err
}

func (r *TaskRepository) Count(req store.CountTasksRequest) (int, error) {
	query := r.DB.Conn.Model((*domain.Task)(nil))
	if req.UserID != nil {
		query.Where("user_id = ?", req.UserID)
	}
	if req.Day != nil {
		query.Where("due_date::date = ?::date", req.Day)
	}
	count, err := query.Count()
	return count, err
}
