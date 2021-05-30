package tasks

type TasksRepository interface {
	Create(req Tasks) (interface{}, error)
	GetList(createDate string) ([]Tasks, error)
	CountTask(userId string, createdDate string) int64
}
