package storages

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/entities"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// LiteDB for working with sqlLite
type liteDB struct {
	db *sql.DB
}

func NewLiteDB() *liteDB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println("liteDB , connect db err:", err)
	} else {
		log.Println("liteDB , connect db success")
		return &liteDB{
			db: db,
		}
	}
	return nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *liteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
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
func (l *liteDB) AddTask(ctx context.Context, t *entities.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *liteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
