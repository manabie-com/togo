package repos

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
)

type ITaskRepo interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) error
}

type TaskRepo struct {
	db *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *TaskRepo) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
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

// AddTask adds a new task to DB
func (l *TaskRepo) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func NewTaskRepo(db *sql.DB) ITaskRepo {
	return &TaskRepo{
		db: db,
	}
}
