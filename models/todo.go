package models

import (
	"time"
)

type Todo struct {
	Id     int       `gorm:"primary_key;auto_increment;not_null"`
	Task   string    `json:"task" binding:"required"`
	Userid int       `json:"userid" binding:"required" gorm:"foreignkey:id ;references:UserId"`
	Date   time.Time `json:"date"`
}
