package model

// User represents the user model
type User struct {
	Base
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"-"`
	Tasks     []*Task `json:"tasks,omitempty"`
}
