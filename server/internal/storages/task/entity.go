package task

// Task reflects tasks in DB
type Task struct {
	ID          uint64
	Content     string
	UserID      uint64
}
