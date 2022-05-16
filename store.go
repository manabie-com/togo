package togo

type Store interface {
	Initialize() error
	CreateUser(username string, password string) (UserEntity, error)
	Login(username string, password string) (int, error)
	CreateTask(taskName string, description string) (TaskEntity, error)
	SetTaskLimit(userId int, taskId int, limit int) error
	GetAllTaskLimitSetting() ([]UserTaskLimitEntity, error)
}

type Cache interface {
	ResetTaskForNewDay(limitSettings []UserTaskLimitEntity) error
	SetTaskLimit(userId int, taskId int, limit int) error
	CheckIfCanDoTask(userId int, taskId int) (int, error)
	RollbackIfCheckFail(userId int, taskId int) error
}
