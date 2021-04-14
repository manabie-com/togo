package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"time"
)

const MaxTodo = 5

type PostgresDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
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
func (l *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}
func (l *PostgresDB) CheckUserLimit(ctx context.Context, userID string) (error, bool) {
	err, user := l.GetUser(ctx, userID)
	if err != nil {
		return err, false
	}

	nowString := time.Now().Format("2006-01-02")

	if user.CurrentDay != nowString {
		err = l.ResetUserLimit(ctx, userID, nowString)
		return err, true
	} else {
		if user.MaxTodo > 0 {
			return nil, true
		}
		return nil, false
	}

}

func (l *PostgresDB) ResetUserLimit(ctx context.Context, userID, currentDay string) error {
	stmt := `UPDATE users SET max_todo = $1, current_day = $2 where id = $3`
	_, err := l.DB.ExecContext(ctx, stmt, MaxTodo, currentDay, userID)
	if err != nil {
		return err
	}

	return nil
}

func (l *PostgresDB) ChangeUserLimit(ctx context.Context, userId string) error {
	stmt := `UPDATE users SET max_todo = max_todo - 1 where id = $1 and max_todo > 0`
	_, err := l.DB.ExecContext(ctx, stmt, userId)
	if err != nil {
		return err
	}

	return nil
}

func (l *PostgresDB) GetUser(ctx context.Context, userID string) (error, storages.User) {
	stmt := "select id, max_todo, current_day from users where id = $1"
	row := l.DB.QueryRowContext(ctx, stmt, userID)

	u := storages.User{}
	err := row.Scan(&u.ID, &u.MaxTodo, &u.CurrentDay)
	if err != nil {
		return err, u
	}

	return nil, u
}

func (l *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := "select id from users where id = $1 and password = $2"
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)

	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
