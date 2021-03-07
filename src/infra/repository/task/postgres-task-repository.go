package task

import (
	"togo/src"
	"togo/src/entity/task"
	"togo/src/infra/client"
	gError "togo/src/infra/error"
)

type TaskRepository struct {
	mapper       *TaskMapper
	db           *client.DB
	errorFactory src.IErrorFactory
}

func (this *TaskRepository) Create(task *task.Task) (*task.Task, error) {
	result := this.db.PConn.Create(&task)

	if result.Error != nil {
		return nil, this.errorFactory.BadRequestError(src.CREATE_TASK_ERROR, result.Error)
	}

	return task, nil
}

func (this *TaskRepository) UpdateById(id string, data *task.Task) (*task.Task, error) {
	return nil, nil
}

func (this *TaskRepository) DeleteById(id string) (bool, error) {
	return true, nil
}

func (this *TaskRepository) FindOne(options interface{}) (*task.Task, error) {
	task := new(task.Task)
	result := this.db.PConn.First(&task, options)

	if result.Error != nil {
		return nil, this.errorFactory.InternalServerError(src.FIND_ONE_TASK_ERROR, result.Error)
	}

	return task, nil
}

func (this *TaskRepository) Find(options interface{}) (*[]task.Task, error) {
	tasks := new([]task.Task)
	result := this.db.PConn.Find(&tasks, options)

	if result.Error != nil {
		return nil, this.errorFactory.InternalServerError(src.FIND_TASK_ERROR, result.Error)
	}

	return tasks, nil
}

func NewTaskRepository() task.ITaskRepository {
	return &TaskRepository{
		&TaskMapper{},
		client.NewDB(),
		gError.NewErrorFactory(),
	}
}
