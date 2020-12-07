package sqllite

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages/task"
)


type taskRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *taskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) Close() error {
	return tr.DB.Close()
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (tr *taskRepository) RetrieveTasks(ctx context.Context, userID int64) ([]task.Task, error) {
	return []task.Task{}, nil
}

// AddTask adds a new task to DB
func (tr *taskRepository) AddTask(ctx context.Context, t task.Task) (int64, error) {
	return 0, nil
}

func (tr *taskRepository) SoftDeleteTask(ctx context.Context, taskId int64) error {
	return nil
}


