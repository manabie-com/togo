package model

type User struct {
	Id     uint   `json:"id" gorm:"id"`
	Name   string `json:"name" gorm:"name"`
	Email  string `json:"email" gorm:"email"`
	Status int    `json:"status" gorm:"status"`
}

func (User) TableName() string {
	return "users"
}
