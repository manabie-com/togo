package usermodel

type UserLimit struct {
	LimitTask int `json:"limit_task" gorm:"column:limit_task" binding:"required"`
}

func (UserLimit) TableName() string {
	return User{}.TableName()
}
