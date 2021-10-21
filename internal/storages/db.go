package storages

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DBModel struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *DBModel) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error) {
	stmt := `SELECT id, content, user_id, created_at FROM tasks WHERE user_id = $1 AND created_at >= $2 AND created_at <= $2`
	rows, err := l.DB.QueryContext(ctx, stmt, userID.String, createdDate.String)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
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
func (l *DBModel) AddTask(ctx context.Context, t *Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match email AND password
func (l *DBModel) ValidateUser(ctx context.Context, email, pwd sql.NullString) bool {
	// stmt := `SELECT id FROM users WHERE email = $1 AND password = $2`
	stmtEmail := `SELECT password FROM users WHERE email = $1`

	row := l.DB.QueryRowContext(ctx, stmtEmail, email.String)
	var u User
	errRow := row.Scan(&u.Password)
	if errRow != nil {
		return false
	}
	hashedPassword := &u.Password
	password := []byte(*hashedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd.String))
	if err != nil {
		return false
	}

	return true
}

//returns one user and error, if any using email query
func (m *DBModel) GetUserFromEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id from users where email = $1
	`
	row := m.DB.QueryRowContext(ctx, query, email)
	var user User
	err := row.Scan(
		&user.ID,
	)
	if err != nil {
		return nil, err
	}

	return &user, err
}
