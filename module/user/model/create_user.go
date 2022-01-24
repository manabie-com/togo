package model

type CreateUser struct {
	Name   *string `json:"name" gorm:"name"`
	Email  *string `json:"email" gorm:"email"`
	Status int     `json:"status" gorm:"status"`
	Id     *uint   `json:"id" gorm:"id"`
}

func (CreateUser) TableName() string {
	return User{}.TableName()
}
