package mock

import (
	"togo/src/entity/task"
)

type TaskRepositoryMock struct {
	GetCreateFunc func(data *task.Task) (*task.Task, error)
	GetCountFunc  func(filter interface{}) (int, error)
}

func (this *TaskRepositoryMock) Create(data *task.Task) (*task.Task, error) {
	return this.GetCreateFunc(data)
}

func (this *TaskRepositoryMock) Count(filter interface{}) (int, error) {
	return this.GetCountFunc(filter)
}

func New_TaskRepository_With_CreateOK_CountEqual5() *TaskRepositoryMock {
	return &TaskRepositoryMock{
		GetCreateFunc: func(data *task.Task) (*task.Task, error) {
			return nil, nil
		},
		GetCountFunc: func(filter interface{}) (int, error) {
			return 5, nil
		},
	}
}

func New_TaskRepository_With_CreateOK_CountLargerThan5() *TaskRepositoryMock {
	return &TaskRepositoryMock{
		GetCreateFunc: func(data *task.Task) (*task.Task, error) {
			return nil, nil
		},
		GetCountFunc: func(filter interface{}) (int, error) {
			return 6, nil
		},
	}
}

func New_TaskRepository_CreateOK_CountLessThan5() *TaskRepositoryMock {
	return &TaskRepositoryMock{
		GetCreateFunc: func(data *task.Task) (*task.Task, error) {
			return &task.Task{Id: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f"}, nil
		},
		GetCountFunc: func(filter interface{}) (int, error) {
			return 4, nil
		},
	}
}

func New_TaskRepository_With_TaskCount_Error() *TaskRepositoryMock {
	return &TaskRepositoryMock{
		GetCreateFunc: func(data *task.Task) (*task.Task, error) {
			return nil, nil
		},
		GetCountFunc: func(filter interface{}) (int, error) {
			return 0, ERROR_500
		},
	}
}
