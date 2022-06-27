package repository

import (
	"gorm.io/gorm"
	"manabieAssignment/internal/core/repository"
	"manabieAssignment/internal/user/repository/dao"
	"time"
)

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) repository.UserRepository {
	return &userRepository{
		gormDB: gormDB,
	}
}

const DateFormat = "01-02-2006"

func (u *userRepository) CountTodosByDay(userID int64, date time.Time) (int64, error) {
	var numOfTodos int64
	err := u.gormDB.Model(&dao.User{}).Where("CAST(created_at as DATE) = ? AND ID = ?", date.Format(DateFormat), userID).Count(&numOfTodos).Error
	return numOfTodos, err
}
