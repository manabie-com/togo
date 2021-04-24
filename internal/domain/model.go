package domain

// Task reflects tasks in DB
type Task struct {
	ID          string
	Content     string
	UserID      string
	CreatedDate string
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}
