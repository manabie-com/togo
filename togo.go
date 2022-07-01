package togo

// User struct represents user fields
type User struct {
	ID            int    `json:"id" db:"id"`
	UserName      string `json:"user_name" db:"user_name"`
	LimitedPerDay int    `json:"limited_per_day" db:"limited_per_day"`
}

// Todo struct represents todo fields
type Todo struct {
	TodoID      int    `json:"todo_id" db:"todo_id"`
	UserID      int    `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
}

// TodoService interface uses for handling Business
type TodoService interface {
	// Adding new todo task to user if not exceed a limited per day
	AddTodoByUser(userName string, t *Todo) (*Todo, error)
}

// TodoDB interface uses for handling Todo Repository
type TodoDB interface {
	// Get user by username
	GetUserByName(userName string) (*User, error)
	// Check if exceed a limited per day
	IsExceedPerDay(u User) (bool, error)
	// Adding new todo task to user if not exceed a limited per day
	AddTodoByUser(u *User, t *Todo, uFlag bool) error
}
