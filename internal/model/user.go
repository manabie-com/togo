package model

// User represents the user model
type User struct {
	Base
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Tasks     []*Task
}
