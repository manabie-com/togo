package storages

type TaskType string

const (
	ACTIVE TaskType = "ACTIVE"
	COMPLETED TaskType = "COMPLETED"
	DELETED TaskType = "DELETED"
)
// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
	Status TaskType `json:"status"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
	MaxTodo int32
}
