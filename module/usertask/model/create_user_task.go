package model

type CreateUserTask struct {
	UserId  *uint `json:"user_id" gorm:"user_id"`
	TaskId *uint `json:"task_id" gorm:"task_id"`
	Status int `json:"status" gorm:"status"`
}

func (CreateUserTask) TableName() string {
	return UserTask{}.TableName()
}
