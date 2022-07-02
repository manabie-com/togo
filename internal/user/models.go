package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Email         string `json:"email"`
	MaxDailyLimit int    `json:"maxDailyLimit"`
}
