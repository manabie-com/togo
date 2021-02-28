package sqlstore

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/task/model"
)

type TaskStore struct {
	*sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (s *TaskStore) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*model.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := s.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		t := &model.Task{}
		if err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (s *TaskStore) AddTask(ctx context.Context, t *model.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)

	return err
}
