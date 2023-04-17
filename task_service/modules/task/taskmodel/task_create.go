package taskmodel

type TaskCreate struct {
	Content string `json:"content" gorm:"column:content;" binding:"required"`
	UserId  int    `json:"user_id" form:"-" gorm:"column:user_id"`
}

func (TaskCreate) TableName() string {
	return Task{}.TableName()
}
