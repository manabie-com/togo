package tasks

type Task struct {
	ID          string `sql:"primary_key" json:"id"`
	UserId      string `sql:"size:100" json:"user_id"`
	Content     string `sql:"size:1000" json:"content"`
	CreatedDate string `sql:"default:CURRENT_TIMESTAMP" json:"created_date"`
}
