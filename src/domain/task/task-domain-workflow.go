package task

import (
	"errors"
	"time"
	"togo/src"
	"togo/src/entity/task"
	"togo/src/entity/user"
	"togo/src/schema"

	"github.com/google/uuid"

	gErrors "togo/src/infra/error"
	taskRepo "togo/src/infra/repository/task"
	userRepo "togo/src/infra/repository/user"
)

type TaskWorkflow struct {
	Repository     task.ITaskRepository
	UserRepository user.IUserRepository
	ErrorFactory   src.IErrorFactory
}

func (this *TaskWorkflow) AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error) {
	tokenData := context.GetTokenData()

	fetchedUser, err := this.UserRepository.FindOne(user.User{ID: tokenData.UserId})
	if err != nil {
		return nil, err
	}

	oldTasksCount, err := this.Repository.Count(task.Task{
		CreatedDate: time.Now().Format("01-02-2006"),
		UserId:      fetchedUser.ID,
	})
	if err != nil {
		return nil, err
	}

	if oldTasksCount >= fetchedUser.MaxTodo {
		return nil, this.ErrorFactory.BadRequestError(src.MAX_TODO_OVER_LIMIT, errors.New("max todo over limit"))
	}

	task := &task.Task{
		Id:          uuid.NewString(),
		Content:     data.Content,
		UserId:      tokenData.UserId,
		CreatedDate: time.Now().Format("01-02-2006"),
	}

	createdTask, err := this.Repository.Create(task)
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
		gErrors.NewErrorFactory(),
	}
}
