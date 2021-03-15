package postgres

import (
	// "context"
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/manabie-com/togo/internal/storages"
)

// PostgresDB for working with posgres
type PostgresDB struct {
	DB *gorm.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PostgresDB) RetrieveTasks(userID, createdDate sql.NullString) ([]*storages.Task, error) {
	// stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	// rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	
	var tasks []storages.Task
	result := l.DB.Where("user_id = ? AND created_date = ?", userID, createdDate).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	// defer rows.Close()

	// var tasks []*storages.Task
	// for rows.Next() {
	// 	t := &storages.Task{}
	// 	err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	tasks = append(tasks, t)
	// }

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	var tks []*storages.Task

	for _, task := range tasks {
		tks = append(tks, &task)
	}

	return tks, nil
}

// CheckMaxTasksPerDay check max task for user
func (l *PostgresDB) CheckMaxTasksPerDay(userID string, maxTasksConfig int16) (bool, int16) {
	var user storages.User
	l.DB.Where("id=?", userID).First(&user);
	
	if  user.CurrentNumberTask >= maxTasksConfig {
		return true, user.CurrentNumberTask
	}
	
	return false, user.CurrentNumberTask
}

// AddTask adds a new task to DB
func (l *PostgresDB) AddTask(t *storages.Task) error {
	// stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	// _, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)

	result := l.DB.Create(&t)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// IncreaseCurrentNumberTask increase +1 when user add new task to DB
func (l *PostgresDB) IncreaseCurrentNumberTask(userID string, curNumTask int16) error {
	result := l.DB.Model(&storages.User{}).Where("id=?", userID).Update("current_number_task", (curNumTask + 1))
	if result.Error != nil {
		fmt.Println("Error Increase Current Number of User!")
		return nil
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *PostgresDB) ValidateUser(userID, pwd sql.NullString) bool {
	// OLD
	// stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	// row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	var user storages.User
	err := l.DB.Where("id = ? AND password = ?", userID, pwd).First(&user).Error

	// NEW
	// u := &storages.User{}
	// err := row.Scan(&u.ID)

	// GORM provides First - it adds LIMIT 1 condition when querying the database, 
	// and it will return error ErrRecordNotFound if no record found.
	if err != nil {
		return false
	}

	return true
}
