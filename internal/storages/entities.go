package storages

// User reflects users data from DB
type TaskStatus struct {
	Code        string
	Name        string
	Description string
}

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	StatusCode  string `json:"status_code"`
	StatusName  string `json:"status_name"`
	DueDate     string `json:"due_date"`
	CreatedDate string `json:"created_date"`
	UpdatedAt   string `json:"updated_at"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
	MaxTodo  int `json:"max_todo"`
}
