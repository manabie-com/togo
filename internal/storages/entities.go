package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	Status      string `json:"status"`
	CreatedDate string `json:"created_date"`
	TargetDate  string `json:"target_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}
