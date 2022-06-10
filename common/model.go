package common

import "time"

type Model struct {
	Id          *uint     `json:"id" gorm:"id"`
	Status      int       `json:"status" gorm:"id"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}
