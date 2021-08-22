package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/util"
)

type (
	TaskUsecase interface {
		ListTasks(ctx context.Context, req *dto.ListTasksRequestDTO) (*dto.ListTaskResponseDTO, error)
		AddTask(ctx context.Context, req *dto.AddTaskRequestDTO) (*dto.AddTaskResponseDTO, error)
	}

	taskUsecase struct {
		store *sqllite.LiteDB
	}
)

func NewTaskUsecase(store *sqllite.LiteDB) TaskUsecase {
	return &taskUsecase{
		store: store,
	}
}

func (u *taskUsecase) ListTasks(ctx context.Context, req *dto.ListTasksRequestDTO) (resp *dto.ListTaskResponseDTO, err error) {
	if req.CreatedDate == "" {
		return nil, errors.New(common.ReasonInvalidArgument.Code())
	}
	_, err = time.Parse("2006-01-02", req.CreatedDate)
	if err != nil {
		return nil, errors.New(common.ReasonDateInvalidFormat.Code())
	}
	tasks, err := u.store.RetrieveTasks(ctx, req.UserID, req.CreatedDate)
	if err != nil {
		log.Printf("retrieve tasks from DB error: %v", err)
		return nil, errors.New(common.ReasonInternalError.Code())
	}

	if len(tasks) == 0 {
		return nil, errors.New(common.ReasonNotFound.Code())
	}

	var listTasks []dto.TaskDTO
	for _, task := range tasks {
		taskDTO := dto.TaskDTO{
			ID:          task.ID.String,
			Content:     task.Content.String,
			UserID:      task.UserID.String,
			CreatedDate: task.CreatedDate.String,
		}
		listTasks = append(listTasks, taskDTO)
	}

	return &dto.ListTaskResponseDTO{
		Data: listTasks,
	}, nil
}

func (u *taskUsecase) AddTask(ctx context.Context, req *dto.AddTaskRequestDTO) (resp *dto.AddTaskResponseDTO, err error) {
	if req.Content == "" {
		log.Println("content is empty")
		return nil, errors.New(common.ReasonInvalidArgument.Code())
	}

	maxTodo, err := u.store.GetMaxTodo(ctx, req.UserID)
	if err != nil {
		log.Printf("get max todo from DB error: %v\n", err)
		return nil, errors.New(common.ReasonInternalError.Code())
	}

	tasks, err := u.store.RetrieveTasks(ctx, req.UserID, time.Now().Format("2006-01-02"))
	if err != nil {
		log.Printf("retrieve tasks from DB error: %v\n", err)
		return nil, errors.New(common.ReasonInternalError.Code())
	}

	if len(tasks) == int(maxTodo) {
		log.Printf("task exceeded limit %v per day\n", maxTodo)
		return nil, errors.New(common.ReasonExceededLimit.Code())
	}

	addTask := storages.Task{
		ID:          util.ConvertSQLNullString(uuid.New().String()),
		UserID:      util.ConvertSQLNullString(req.UserID),
		Content:     util.ConvertSQLNullString(req.Content),
		CreatedDate: util.ConvertSQLNullString(time.Now().Format("2006-01-02")),
	}

	err = u.store.AddTask(ctx, addTask)
	if err != nil {
		log.Printf("Add task DB error: %s\n", err)
		return nil, errors.New(common.ReasonInternalError.Code())
	}
	return &dto.AddTaskResponseDTO{
		Data: dto.TaskDTO{
			ID:          addTask.ID.String,
			UserID:      addTask.UserID.String,
			Content:     addTask.Content.String,
			CreatedDate: addTask.CreatedDate.String,
		},
	}, nil
}
