package model

import "gorm.io/gorm"

type GormTable interface {
	TableName() string
}

type UserModel struct {
	// ID               uint `gorm:"primaryKey,autoIncrement"`
	gorm.Model
	Fullname         string
	DailyRecordLimit uint
}

type UserRedisModel struct {
	ID               uint
	CurrentUsage     uint
	DailyRecordLimit uint
}

func (UserModel) TableName() string {
	return "users"
}

type TodoItemModel struct {
	// ID      uint `gorm:"primaryKey,autoIncrement"`
	gorm.Model
	Content string
	IsDone  bool
	UserID  uint
}

func (TodoItemModel) TableName() string {
	return "user_todo_items"
}
