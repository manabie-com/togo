package models

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	UserID          string    `gorm:"type:varchar(256);not null;unique"`
	DailyTasksLimit uint      `gorm:"default:16"`
	MaxDailyTasks   uint      `gorm:"default:16"`
	LastUpdatedTask time.Time `gorm:"not null"`
	TasksList       []Task    `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (um *UserModel) CreateUser(ctx context.Context, userId string, maxTasks uint) (newUser *User, err error) {
	newUser = &User{
		UserID:          userId,
		DailyTasksLimit: maxTasks,
		MaxDailyTasks:   maxTasks,
		LastUpdatedTask: time.Now(),
	}
	if err := um.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (um *UserModel) UpdateUser(ctx context.Context, user *User) (updated bool, err error) {
	if err = um.db.Model(&user).Select("*").Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func (um *UserModel) GetUserByUserId(ctx context.Context, userId string) (user *User, err error) {
	if err = um.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
