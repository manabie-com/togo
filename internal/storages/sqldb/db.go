package sqldb

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

// SqlDB implement storages.Repositoty for sql dbs
type SqlDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks by userID AND createDate.
func (repo *SqlDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := repo.DB.QueryContext(ctx, stmt, toSqlNullString(userID), toSqlNullString(createdDate))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*storages.Task{}
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
func (repo *SqlDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := repo.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks by userID AND password
func (repo *SqlDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := repo.DB.QueryRowContext(ctx, stmt, toSqlNullString(userID), toSqlNullString(pwd))
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

// MaxTodo returns a user's max_todo (max daily tasks create requests) by userID
func (repo *SqlDB) MaxTodo(ctx context.Context, userID string) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = $1`
	row := repo.DB.QueryRowContext(ctx, stmt, toSqlNullString(userID))
	u := &storages.User{}
	err := row.Scan(&u.MaxTodo)
	if err != nil {
		return 0, err
	}

	return u.MaxTodo, nil
}

func (repo *SqlDB) LoadTasksCount(ctx context.Context, userID, createdDate string) (cnt int, err error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2`
	err = repo.DB.QueryRowContext(ctx, stmt, toSqlNullString(userID), toSqlNullString(createdDate)).Scan(&cnt)
	return
}

func toSqlNullString(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}
