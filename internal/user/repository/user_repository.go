package repository

import (
	"errors"
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

func (u *userRepository) IsUserHavingMaxTodo(userID int64, date time.Time) error {
	var user dao.User
	countTodosQuery := u.gormDB.Select("COUNT(*)").Where("CAST(created_at as DATE) = ? AND user_id = ?", date.Format(DateFormat), userID).Table("todos")
	rows := u.gormDB.Where("id = ? AND max_todo > (?)", userID, countTodosQuery).Take(&user).RowsAffected
	if rows == 0 {
		return errors.New("user have too many todos")
	}
	return nil
}
