package tasks

type Tasks struct {
	ID          string `json:"id" form:"id"`
	Content     string `json:"content" form:"content"`
	UserId      string `json:"user_id" form:"user_id"`
	CreatedDate string `json:"created_date" form:"created_date"`
}

func (Tasks) TableName() string {
	return "tasks"
}
