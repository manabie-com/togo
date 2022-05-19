package togo

type TaskLimiterService interface {
	CreateTask(rq TaskCreationRequest) (TaskCreationResponse, error)
	DoTask(userId int, taskId int) (string, error)
	SetTaskLimit(userId int, taskId int, limit int) (string, error)
	ResetDailyTask() (string, error)
}

type UserCrudService interface {
	CreateUser(rq UserRequest) (UserResponse, error)
	Login(rq UserRequest) (string, error)
}
