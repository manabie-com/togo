package sqllite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/internal/storages"
	"golang.org/x/crypto/bcrypt"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

func (l *LiteDB) InitTables() error {

	stmt := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT NOT NULL,
		password TEXT NOT NULL,
		max_todo INTEGER DEFAULT 5 NOT NULL,
		CONSTRAINT users_PK PRIMARY KEY (id)
	);
	`
	_, err := l.DB.Exec(stmt)
	if err != nil {
		return err
	}

	stmt2 := `CREATE TABLE IF NOT EXISTS tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = l.DB.Exec(stmt2)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
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
	ok, err := l.CanUserCreateTodo(ctx, t)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("user %s cannot create any more todos for %s", t.UserID, t.CreatedDate)
	}

	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err = l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// AddUser adds a new user to DB
func (l *LiteDB) AddUser(ctx context.Context, user *storages.User) error {
	stmt := `INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = l.DB.ExecContext(ctx, stmt, user.ID, string(hash), user.MaxTodo)
	if err != nil {
		return err
	}

	return nil
}

// MaxTodo checks the user account for how many max todos it has
func (l *LiteDB) MaxTodo(ctx context.Context, userID string) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	var maxTodo int
	err := row.Scan(&maxTodo)
	if err != nil {
		return 0, err
	}

	return maxTodo, nil
}

// CanUserCreateTodo checks if the user can create a todo.
// this will return false if the user has no more todos left for the day
func (l *LiteDB) CanUserCreateTodo(ctx context.Context, t *storages.Task) (bool, error) {
	stmt := `SELECT count(id) FROM tasks where user_id = ? AND created_date = ?`
	row := l.DB.QueryRowContext(ctx, stmt, t.UserID, t.CreatedDate)

	num := 0
	err := row.Scan(&num)
	if err != nil {
		return false, err
	}
	max, err := l.MaxTodo(ctx, t.UserID)
	if err != nil {
		return false, err
	}

	return num < max, nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT password FROM users WHERE id = ?`

	row := l.DB.QueryRowContext(ctx, stmt, userID)

	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		// should return error here
		fmt.Println(passwordHash, err)
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(pwd))
	if err != nil {
		return false
	}

	return true
}
