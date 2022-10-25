package user

import "time"

type User struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"type:varchar(20);unique;not null"`
	Password     string    `json:"password" gorm:"type:varchar(72);not null"`
	MaxDailyTask int       `json:"max_daily_task" gorm:"type:int;default:5"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (m User) TableName() string {
	return "user"
}
