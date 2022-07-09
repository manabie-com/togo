package togo

import (
	"context"
)

// General errors.
const (
	ErrSuccessAddingNewTodo  = Error("Successfully adding new todo task")
	ErrIsExceedLimitedPerDay = Error("You have reached the limit of adding todo task per day")
	ErrPageNotFound          = Error("Your page cannot be found")
	ErrHTTPMethodNotAllowed  = Error("HTTP Method is not allowed")
	ErrCouldNotParseObject   = Error("Could not parse Todo object")
)

// Error is a domain error encountered while processing chronograf requests
type Error string

// Error ToString method
func (e Error) Error() string {
	return string(e)
}

// Todo represents a todo in the system. Todos are typically created via API.
type Todo struct {
	ID int `json:"id"`
	// Description of todo task
	Description string `json:"description"`

	// Timestamps for todo creation & last update.
	// CreatedAt time.Time `json:"created_at"`
	// CreatedBy string    `json:"created_by"`
	// UpdatedAt time.Time `json:"updated_at"`
	// UpdatedBy string    `json:"updated_by"`
}

// TodoService represents a service for managing todos.
type TodoService interface {
	// Add new todo task
	Add(ctx context.Context, t *Todo, username string) (*Todo, error)
}

// TodoRepo represents a repository for managing todos.
type TodoRepo interface {
	// Add new todo task
	Add(ctx context.Context, t *Todo, u *User) (*Todo, error)
	// Add new todo task with new user
	AddWithNewUser(ctx context.Context, t *Todo, u *User) (*Todo, error)
}

// User represents a user in the system.
type User struct {
	ID int `json:"id"`
	// User's preferred name & limited todo per day.
	Username      string `json:"user_name"`
	LimitedPerDay int    `json:"limited_per_day"`

	// Timestamps for user creation & last update.
	// CreatedAt time.Time `json:"created_at"`
	// CreatedBy string    `json:"created_by"`
	// UpdatedAt time.Time `json:"updated_at"`
	// UpdatedBy string    `json:"updated_by"`
}

// UserService represents a service for managing users.
type UserService interface {
	// Retrieves a user by name along with their associated auth objects.
	GetUserByName(ctx context.Context, username string) (*User, error)
	// Check the number of todos created has reached the limited per day or not.
	IsExceedPerDay(ctx context.Context, u *User) (bool, error)
}

// UserRepo represents a repository for managing users.
type UserRepo interface {
	// Retrieves a user by name along with their associated auth objects.
	GetUserByName(ctx context.Context, username string) (*User, error)
	// Check the number of todos created has reached the limited per day or not.
	IsExceedPerDay(ctx context.Context, u *User) (bool, error)
}

// ErrResponse represent a error with json format
type TemplateResponse struct {
	// HTTP response status codes
	Status int
	// Error message
	Message string
	// Todo object
	Data Todo
}
