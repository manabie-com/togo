package repositories

import (
	"gorm.io/gorm"
)

const (
	UsersTableName = "users"
	fieldUserIDName = "id"

)

func (User) TableName() string {
	return UsersTableName
}


type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) ValidateUser(userID, password string) bool {
	var user  User
	err := u.db.Where("id = ? and password = ?", userID, password).First(&user).Error
	if err != nil {
		return false
	}
	return true
}