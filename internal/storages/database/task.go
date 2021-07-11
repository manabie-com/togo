package database

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entity"
)

// repository for working with sqllite
type taskRepository struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *taskRepository) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entity.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		t := &entity.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *taskRepository) AddTask(ctx context.Context, t *entity.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}
