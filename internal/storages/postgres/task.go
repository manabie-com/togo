package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

const (
	RetrieveTasksStmt = `SELECT id, content, user_id, created_date, number_in_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	AddTaskStmt       = `INSERT INTO tasks (id, content, user_id, created_date, number) VALUES ($1, $2, $3, $4, $5)`
)

type taskStore struct {
	DB *sql.DB
}

func (s *taskStore) RetrieveTasks(ctx context.Context, task *storages.Task) ([]*storages.Task, error) {
	rows, err := s.DB.QueryContext(ctx, RetrieveTasksStmt, task.UserID, task.CreatedDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate, &t.NumberInDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskStore) AddTask(ctx context.Context, t *storages.Task) error {
	// create id
	t.ID = uuid.New().String()

	_, err := s.DB.ExecContext(ctx, AddTaskStmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate, &t.NumberInDate)
	return err
}

func NewTaskStore(db *sql.DB) storages.TaskStore {
	return &taskStore{
		DB: db,
	}
}
