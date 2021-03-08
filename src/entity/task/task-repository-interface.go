package task

type ITaskRepository interface {
	Create(task *Task) (*Task, error)
	Count(filter interface{}) (int, error)
}
