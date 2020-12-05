package task

// Task reflects tasks in DB
type Task struct {
	ID          string
	Content     string
	UserID      string
	CreatedDate string
}
