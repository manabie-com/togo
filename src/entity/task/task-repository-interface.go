package task

type ITaskRepository interface {
	Create(task *Task) (*Task, error)
	FindOne(options interface{}) (*Task, error)
	Find(options interface{}) (*[]Task, error)
	UpdateById(id string, data *Task) (*Task, error)
	DeleteById(id string) (bool, error)
}
