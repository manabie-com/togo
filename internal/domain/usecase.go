package domain

type TaskUseCase interface {
	AddTask(Task) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type AuthUseCase interface {
	FindUserByID(ID string) (User, error)
	ValidateUserPassword(given string, hashed string) bool
	CreateUser(ID string, password string) error
}
