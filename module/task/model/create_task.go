package model

type CreateTask struct {
	Id     *uint   `json:"id" gorm:"id"`
	Name   *string `json:"name" gorm"name"`
	Status int     `json:"status" gorm:"status"`
}

func (CreateTask) TableName() string {
	return Task{}.TableName()
}
