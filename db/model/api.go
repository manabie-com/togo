package model

// DatabaseAPI interface
type DatabaseAPI interface {
	AddUser(userID string, fields map[string]interface{}) (*User, error)

	UpdateUser(userID string, fields map[string]interface{}) (*User, error)

	CreateUser(user User) error

	DeleteUser(user User) error

	GetUsers() []User

	GetUser(userID string) (*User, error)

	GetUserByName(userName string) (*User, error)

	IsUserNotExists(userName string) bool

	GetListUserID() []string

	CreateTask(userID, taskName string) error

	Close() error
}
