package services

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (srv *TaskService) CreateTasks() (map[string]interface{}, error) {
	return map[string]interface{}{
		"info": []map[string]interface{}{
			{
				"Title":       "Task 1",
				"Description": "My first task",
			},
			{
				"Title":       "Task 2",
				"Description": "My second task",
			},
		},
	}, nil
}
