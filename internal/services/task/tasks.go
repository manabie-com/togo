package task

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/entity"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/pkg/customcontext"
)

var (
	nowFunc           = time.Now
	timeLayout = "2006-01-02"
)

type (
	service struct {
		taskStorage storages.TaskStorage
		userStorage storages.UserStorage
	}
	Service interface {
		List(ctx context.Context, createdDate string) ([]*entity.Task, error)
		Create(ctx context.Context, content string) (*entity.Task, error)
	}
)

func NewTaskService(taskStorage storages.TaskStorage, userStorage storages.UserStorage) Service {
	return &service{
		taskStorage: taskStorage,
		userStorage: userStorage,
	}
}

func (s *service) List(ctx context.Context, createdDate string) ([]*entity.Task, error) {
	userID := customcontext.UserIDFromCtx(ctx)
	return s.taskStorage.RetrieveTasks(ctx, userID, createdDate)
}

func (s *service) Create(ctx context.Context, content string) (*entity.Task, error) {
	userID := customcontext.UserIDFromCtx(ctx)

	reached, err := s.reachedOutTaskTodoLimit(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !reached {
		return nil, ErrReachedOutTaskTodoPerDay
	}
	t := &entity.Task{}
	t.UserID = userID
	t.Content = content
	return t, s.taskStorage.AddTask(ctx, t)
}

func (s *service) reachedOutTaskTodoLimit(ctx context.Context, userID string) (bool, error) {

	// should ignore error due to userID is passed from ctx after passing auth middleware.
	user, _ := s.userStorage.FindByID(ctx, userID)
	if user == nil {
		return false, NewUserNotExistedError(userID)
	}
	now := nowFunc()
	today := now.Format(timeLayout)
	nTaskCreated, err := s.taskStorage.GetNumberOfTasks(ctx, userID, today)
	if err != nil {
		return false, err
	}
	if nTaskCreated > user.MaxTodoPerday {
		return true, nil
	}
	return false, nil
}
