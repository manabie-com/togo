package task

type Task struct {
	Id          string `gorm:"column:id"`
	Content     string `gorm:"column:content"`
	UserId      string `gorm:"column:user_id"`
	CreatedDate string `gorm:"column:created_date"`
}
