package model

type Task struct {
	Id     uint32 `json:"id" gorm:"id"`
	Name   string `json:"name" gorm"name"`
	Status int    `json:"status" gorm:"status"`
}

func (Task) TableName() string {
	return "tasks"
}
