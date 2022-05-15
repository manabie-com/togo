package model

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/pkg/database"
)

// user table users
type user struct {
	ID              int16  `gorm:"column:id;primaryKey"`
	Name            string `gorm:"column:name"`
	LimitTaskPerDay int16  `gorm:"column:limit_task_per_day"`
	IsDeleted       bool   `gorm:"column:is_deleted"`
	Tasks           []task
}

// NewUser user constructor
func NewUser() *user {
	return &user{}
}

// filter filter user is available
func (*user) filter() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("users.is_deleted IS FALSE")
	}
}

// GetUserTasksByUserID get user and their tasks today
func (*user) GetUserTasksTodayByUserID(userID int16) (*user, error) {
	user := user{}
	task := task{}
	err := database.DB().
		Scopes(
			where("id = ?", userID),
			user.filter(),
			preload(
				"Tasks",
				task.filter(),
				task.assignToDay(),
			),
		).
		Take(&user).Error
	return &user, err
}
