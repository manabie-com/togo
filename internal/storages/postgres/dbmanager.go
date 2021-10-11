package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cuongtop4598/togo-interview/togo/internal/helper"
	"github.com/cuongtop4598/togo-interview/togo/internal/storages"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type TodoRepository interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

type DBmanager struct {
	*gorm.DB
}

func NewDBManager() (TodoRepository, error) {
	db, err := NewGormDB()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		storages.User{},
		storages.Task{},
	)
	if err != nil {
		return nil, err
	}
	return &DBmanager{db}, nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (m *DBmanager) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	tasks := []*storages.Task{}
	if userID.Valid && createdDate.Valid {
		if err := m.Where(&storages.Task{
			UserID:      userID.String,
			CreatedDate: createdDate.String,
		}).Find(&tasks).Error; err != nil {
			return nil, err
		}
		return tasks, nil
	}
	return nil, errors.New("userID and createdDate is not null")
}

// AddTask adds a new task to DB
func (m *DBmanager) AddTask(ctx context.Context, t *storages.Task) error {
	tasks := []*storages.Task{}
	if err := m.Where("created_at >= ?", now.BeginningOfDay()).Find(&tasks).Error; err != nil {
		return err
	}
	userInfo := storages.User{}
	if err := m.Where(&storages.User{ID: t.UserID}).Find(&userInfo).Error; err != nil {
		return err
	}
	if len(tasks) < userInfo.MaxTodo {
		task := &storages.Task{
			ID:          uuid.New(),
			Content:     t.Content,
			UserID:      t.UserID,
			CreatedDate: t.CreatedDate,
		}
		if err := m.Create(task).Error; err != nil {
			return err
		}
		return nil
	} else {
		return helper.ErrExceedMaxTaskPerDay
	}
}

// ValidateUser returns tasks if match userID AND password
func (m *DBmanager) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	user := &storages.User{}
	if userID.Valid && pwd.Valid {
		if err := m.Where(&storages.User{
			ID:       userID.String,
			Password: pwd.String,
		}).First(&user).Error; err != nil {
			return false
		}
		return true
	}
	return false
}
