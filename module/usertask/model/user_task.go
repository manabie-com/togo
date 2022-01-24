package model

type UserTask struct {
	UserId uint `json:"user_id" gorm:"user_id"`
	TaskId uint `json:"task_id" gorm:"task_id"`
	Status int  `json:"status" gorm:"status"`
}

func (UserTask) TableName() string {
	return "user_tasks"
}
