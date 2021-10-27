package repository

// import (
// 	"context"
// 	"database/sql"

// 	"github.com/manabie-com/togo/internal/core/domain"
// )

// // LiteDB for working with sqllite
// type LiteDB struct {
// 	DB *sql.DB
// }

// // RetrieveTasks returns tasks if match userId AND createDate.
// func (l *LiteDB) RetrieveTasks(ctx context.Context, userId, createdDate sql.NullString) ([]*domain.Task, error) {
// 	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
// 	rows, err := l.DB.QueryContext(ctx, stmt, userId, createdDate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var tasks []*domain.Task
// 	for rows.Next() {
// 		t := &domain.Task{}
// 		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, t)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return tasks, nil
// }

// // AddTask adds a new task to DB
// func (l *LiteDB) AddTask(ctx context.Context, t *domain.Task) error {
// 	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
// 	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // ValidateUser returns tasks if match userId AND password
// func (l *LiteDB) ValidateUser(ctx context.Context, userId, pwd sql.NullString) bool {
// 	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
// 	row := l.DB.QueryRowContext(ctx, stmt, userId, pwd)
// 	u := &domain.User{}
// 	err := row.Scan(&u.ID)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
