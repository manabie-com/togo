package task

// Task reflects tasks in DB
type Task struct {
	ID          int64
	Content     string
	UserID      int64
}
