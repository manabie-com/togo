package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}
