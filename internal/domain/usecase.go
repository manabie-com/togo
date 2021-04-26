package domain

type TaskUseCase interface {
	AddTask(Task) error
	GetTasksByUserIDAndDate(userID string, date string) ([]Task, error)
}

type AuthUseCase interface {
	ValidateUser(userID, rawPassword string) (bool, error)
	// ValidateUserPassword(given string, hashed string) bool
	CreateUser(ID string, Password string) error
}
