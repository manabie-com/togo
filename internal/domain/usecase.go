package domain

type TaskUseCase interface {
	AddTask(Task) error
	GetTasksByUserIDAndDate(userID string, date string) ([]Task, error)
}

type AuthUseCase interface {
	FindUserByID(ID string) (User, error)
	ValidateUserPassword(given string, hashed string) bool
	CreateUser(ID string, Password string) error
}
