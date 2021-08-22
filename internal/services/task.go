package services

import (
	"context"
	"github.com/google/uuid"
	taskApi "github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/tools"
	"time"
)

type TaskService struct {
	repo        storages.ITaskRepo
	contextTool tools.IContextTool
}

func (ts *TaskService) ListTasksByUserAndDate(ctx context.Context, request taskApi.ListTaskRequest) (*taskApi.ListTaskResponse, *tools.TodoError) {
	id, err := ts.contextTool.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	taskEntities, err := ts.repo.RetrieveTasksStore(ctx, storages.RetrieveTasksParams{UserID: id, CreatedDate: request.CreatedDate})
	if err != nil {
		return nil, err
	}
	var taskRes []taskApi.Task
	for _, entity := range taskEntities {
		taskRes = append(taskRes, taskApi.Task{
			ID:          entity.ID,
			Content:     entity.Content,
			UserID:      entity.UserID,
			CreatedDate: entity.CreatedDate,
		})
	}
	return &taskApi.ListTaskResponse{Data: taskRes}, nil
}

func (ts *TaskService) AddTask(ctx context.Context, request taskApi.AddTaskRequest) (*taskApi.AddTaskResponse, *tools.TodoError) {
	id, err := ts.contextTool.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	taskEntity := storages.AddTaskParams{
		Content:     request.Content,
		UserID:      id,
		CreatedDate: now.Format("2006-01-02"),
		ID:          uuid.New().String(),
	}
	err = ts.repo.AddTaskStore(ctx, taskEntity)
	if err != nil {
		return nil, err
	}
	return &taskApi.AddTaskResponse{
		Data: taskApi.Task{
			Content:     taskEntity.Content,
			UserID:      taskEntity.UserID,
			CreatedDate: taskEntity.CreatedDate,
			ID:          taskEntity.ID,
		},
	}, nil
}

func NewTaskService(repo storages.ITaskRepo, contextTool tools.IContextTool) taskApi.ITaskService {
	return &TaskService{
		repo:        repo,
		contextTool: contextTool,
	}
}
