package task

import (
	"time"
	"togo/src"
	"togo/src/entity/task"
	"togo/src/entity/user"
	"togo/src/schema"

	"github.com/google/uuid"

	taskRepo "togo/src/infra/repository/task"
	userRepo "togo/src/infra/repository/user"
)

type TaskWorkflow struct {
	repository     task.ITaskRepository
	userRepository user.IUserRepository
}

func (this *TaskWorkflow) AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error) {
	tokenData := context.GetTokenData()

	if _, err := this.userRepository.FindOne(user.User{ID: tokenData.UserId}); err != nil {
		return nil, err
	}

	task := &task.Task{
		Id:          uuid.NewString(),
		Content:     data.Content,
		UserId:      tokenData.UserId,
		CreatedDate: time.Now(),
	}

	createdTask, err := this.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return &schema.AddTaskResponse{
		TaskId: createdTask.Id,
	}, nil
}

func NewTaskWorkflow() ITaskWorkflow {
	return &TaskWorkflow{
		taskRepo.NewTaskRepository(),
		userRepo.NewUserRepository(),
	}
}
