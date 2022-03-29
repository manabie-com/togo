package models

import (
	"fmt"
	"time"
	"togo/globals/database"
)

type User struct {
	ID			uint	`gorm:"primaryKey;autoIncrement:false" json:"id"`
	Tasks 		[]Task	`gorm:"foreignKey:UserID" json:"user_tasks"`
	TasksPerDay uint8	`gorm:"default:8" json:"tasks_per_day"`
	CreatedAt		*time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt		*time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func UserById(uid uint) (*User, error) {
	user := User{}
	err := database.SQL.First(&user, uid).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func CreateUserWithTask(uid uint, taskDetail string) (*User, error) {
	user := User{ ID: uid, Tasks: []Task{ {Detail: taskDetail} }}
	err := database.SQL.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

func UserAtDailyLimit(uid uint) (bool, error) {
	currentDate := time.Now().Format("2006-01-02")
	startTime := fmt.Sprintf("%sT00:00:00Z", currentDate)
	endTime := fmt.Sprintf("%sT23:59:59Z", currentDate)
	user := User{ID: uid}
	err := database.SQL.Preload("Tasks", "created_at >= ? AND created_at <= ?", startTime, endTime).Find(&user).Error
	if err != nil {
		return false, err
	}
	return uint8(len(user.Tasks)) >= user.TasksPerDay, nil
}