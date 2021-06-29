package repositories

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	usersTableName = "users"
	fieldUserIDName = "id"
	fieldPasswordName = "password"
)

func (User) TableName() string {
	return usersTableName
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}

type UserRepo interface {
	ValidateUser(userID, password string) bool
	GetMaxToDoOfUser(userID string) (int, error)
}

func (u *UserRepoImpl) ValidateUser(userID, password string) bool {
	var user  User
	err := u.db.Where(fmt.Sprintf("%s = ? and %s = ?",fieldUserIDName, fieldPasswordName), userID, password).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func (u *UserRepoImpl) GetMaxToDoOfUser(userID string) (int, error) {
	var user User
	err := u.db.Where(fmt.Sprintf("%s = ?", fieldUserIDName), userID).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.MaxTodo, nil
}