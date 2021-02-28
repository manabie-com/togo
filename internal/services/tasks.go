package services

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/services/helper"
	tasksqlstore "github.com/manabie-com/togo/internal/storages/task/sqlstore"
	usersqlstore "github.com/manabie-com/togo/internal/storages/user/sqlstore"
	"github.com/manabie-com/togo/pkg/common/xerrors"
	"github.com/manabie-com/togo/up"
	"time"

	"github.com/google/uuid"
	taskmodel "github.com/manabie-com/togo/internal/storages/task/model"
)

var _ up.TaskService = &TaskService{}

type TaskService struct {
	userstore *usersqlstore.UserStore
	taskstore *tasksqlstore.TaskStore

	mapUserAndTodos map[string]*helper.ToDos
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{
		userstore:       usersqlstore.NewUserStore(db),
		taskstore:       tasksqlstore.NewTaskStore(db),
		mapUserAndTodos: make(map[string]*helper.ToDos),
	}
}

func (s *TaskService) ListTasks(ctx context.Context, req *up.ListTasksRequest) (resp *up.ListTasksResponse, _ error) {
	id, _ := userIDFromCtx(ctx)
	tasks, err := s.taskstore.RetrieveTasks(
		ctx,
		sql.NullString{String: id, Valid: true},
		sql.NullString{String: req.CreatedDate, Valid: true},
	)
	if err != nil {
		return nil, xerrors.Error(xerrors.Internal, err)
	}

	var tasksResp up.ListTasksResponse
	for _, task := range tasks {
		tasksResp = append(tasksResp, &up.Task{
			ID:          task.ID,
			Content:     task.Content,
			UserID:      task.UserID,
			CreatedDate: task.CreatedDate,
		})
	}

	return &tasksResp, nil
}

func (s *TaskService) AddTask(ctx context.Context, req *up.AddTaskRequest) (resp *up.Task, _ error) {
	userID, _ := userIDFromCtx(ctx)
	now := time.Now()

	task := &taskmodel.Task{
		ID:          uuid.New().String(),
		Content:     req.Content,
		UserID:      userID,
		CreatedDate: now.Format("2006-01-02"),
	}

	{
		todos, ok := s.mapUserAndTodos[userID]
		if !ok {
			user, err := s.userstore.FindByID(ctx, sql.NullString{String: userID, Valid: true})
			if err != nil {
				return nil, xerrors.Error(xerrors.Internal, err)
			}

			numOfTasks, err := s.taskstore.CountByUserID(
				ctx,
				sql.NullString{String: userID, Valid: true},
				sql.NullString{String: task.CreatedDate, Valid: true})
			if err != nil {
				return nil, xerrors.Error(xerrors.Internal, err)
			}

			todos = helper.NewToDos(user.MaxTodo, numOfTasks)
			s.mapUserAndTodos[userID] = todos
		}

		if !todos.CanAddNewTodo() {
			return nil, xerrors.ErrorM(xerrors.InvalidArgument, nil, "the number of todos daily limit is reached")
		}
	}

	err := s.taskstore.AddTask(ctx, task)
	if err != nil {
		return nil, xerrors.Error(xerrors.Internal, err)
	}

	return &up.Task{
		ID:          task.ID,
		Content:     task.Content,
		UserID:      task.UserID,
		CreatedDate: task.CreatedDate,
	}, nil
}
