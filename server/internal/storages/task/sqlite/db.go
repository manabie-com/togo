package sqllite

import (
	"context"
	"database/sql"
	taskUsecase "github.com/HoangVyDuong/togo/internal/usecase/task"

	"github.com/HoangVyDuong/togo/internal/storages/task"
)


type TaskRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) taskUsecase.Repository {
	return &TaskRepository{db}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *TaskRepository) RetrieveTasks(ctx context.Context, userID int64) ([]task.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? `
	rows, err := l.DB.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		t := &task.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return []task.Task{}, nil
}

// AddTask adds a new task to DB
func (l *TaskRepository) AddTask(ctx context.Context, t task.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}


func (l *TaskRepository) Delete(ctx context.Context, taskId int64) bool {
	return false
}
