package postgres

import (
	"errors"
	"time"

	"github.com/manabie-com/togo/pkg/model"

	"github.com/jinzhu/gorm"
)

type TaskFilter struct {
	UserId      int
	CreatedDate *time.Time
}

type Repository interface {
	// user
	GetUser(userName string) (*model.User, error)
	SaveUser(user *model.User) error

	// task
	FindTask(filter TaskFilter) ([]model.Task, error)
	SaveTask(task *model.Task) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	if db == nil {
		return nil, errors.New("db connection is nil")
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) GetUser(userName string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) SaveUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *repository) FindTask(filter TaskFilter) ([]model.Task, error) {
	var result []model.Task
	builder := r.db.Model(&model.Task{})
	if filter.UserId != 0 {
		builder = builder.Where("user_id = ?", filter.UserId)
	}
	if filter.CreatedDate != nil {
		builder = builder.Where("date(created_date) = ?", filter.CreatedDate.Format("2006-01-02"))
	}

	if err := builder.Order("created_date").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}

func (r *repository) SaveTask(task *model.Task) error {
	return r.db.Save(task).Error

}
