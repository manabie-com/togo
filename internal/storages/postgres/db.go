package postgres

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *Storage) RetrieveTasks(userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	var tasks []*storages.Task

	if err := l.DB.Raw(stmt, userID, createdDate).Scan(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *Storage) AddTask(t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	if err := l.DB.Exec(stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate).Error; err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *Storage) ValidateUser(userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	user := &storages.User{}
	if err := l.DB.Raw(stmt, userID, pwd).Scan(user).Error; err != nil {
		return false
	}

	return true
}