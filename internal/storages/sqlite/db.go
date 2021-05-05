package sqllite

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) CreateDatabase() error {
	if nil != l.DB {
		_, err := l.DB.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id TEXT NOT NULL,
				password TEXT NOT NULL,
				max_todo INTEGER DEFAULT 5 NOT NULL,
				CONSTRAINT users_PK PRIMARY KEY (id)
			);

			INSERT INTO users(id, password, max_todo)
				SELECT 'firstUser', 'example', 5
			WHERE NOT EXISTS (
				SELECT 1 FROM users WHERE id = 'firstUser' AND password = 'example'
			);		
			
			CREATE TABLE IF NOT EXISTS tasks (
				id TEXT NOT NULL,
				content TEXT NOT NULL,
				user_id TEXT NOT NULL,
				created_date TEXT NOT NULL,
				CONSTRAINT tasks_PK PRIMARY KEY (id),
				CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`)
		return err
	}
	return nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
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
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

// ValidateTask returns int if match date
func (l *LiteDB) ValidateTask(ctx context.Context, now time.Time) int {
	stmt := `SELECT count(id) FROM tasks WHERE created_date = $1`
	row := l.DB.QueryRowContext(ctx, stmt, now.Format("2006-01-02"))
	countTask := 0
	err := row.Scan(&countTask)
	if err != nil {
		return countTask
	}

	return countTask
}
