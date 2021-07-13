package entity

// Task represents for Task entity in system.
type Task struct {
	ID          string
	Content     string
	UserID      string
	CreatedDate string
}

// User represents for User entity in system.
type User struct {
	ID            string
	Password      string
	MaxTodoPerday int
}
