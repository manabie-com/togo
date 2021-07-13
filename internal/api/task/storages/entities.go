package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content" validate:"required"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
