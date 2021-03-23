package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/entities"
	"log"
	"time"
)

type taskRepository struct {
}

type userRepository struct {
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *taskRepository) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entities.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := db.QueryContext(ctx, stmt, generateToNullStringValue(userID), generateToNullStringValue(createdDate))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entities.Task
	for rows.Next() {
		t := &entities.Task{}
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
func (l *taskRepository) AddTask(ctx context.Context, t *entities.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *userRepository) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = $1 and password = $2`
	rows, err := db.QueryContext(ctx, stmt, userID, pwd)
	defer rows.Close()

	if err != nil {
		return false
	}
	u := &entities.User{}

	rows.Next()
	err = rows.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

func generateToNullStringValue(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func (l *taskRepository) CountTaskPerDayByUserID(ctx context.Context, userID string) (uint, error) {
	stmt := `SELECT COUNT(id) FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := db.QueryContext(ctx, stmt, userID, time.Now().Format("2006-01-02"))
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var count uint
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count, err
}
