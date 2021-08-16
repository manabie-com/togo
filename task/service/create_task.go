package service

import (
	"context"
	"github.com/manabie-com/togo/shared"
	"github.com/manabie-com/togo/task/dto"
	taskmodel "github.com/manabie-com/togo/task/model"
	usermodel "github.com/manabie-com/togo/user/model"
	"time"
)

type UserStorage interface {
	FindById(_ context.Context, id int) (*usermodel.User, error)
}

type TaskStorage interface {
	Create(ctx context.Context, data *taskmodel.Task) error
	Count(ctx context.Context, conditions map[string]interface{}) (*int64, error)
}

type createTaskService struct {
	uStore    UserStorage
	tStore    TaskStorage
	requester shared.Requester
}

func NewCreateTaskService(uStore UserStorage, tStore TaskStorage, requester shared.Requester) *createTaskService {
	return &createTaskService{
		uStore:    uStore,
		tStore:    tStore,
		requester: requester,
	}
}

func (s *createTaskService) CreateTask(ctx context.Context, data *dto.CreateTaskRequest) error {
	user, err := s.uStore.FindById(ctx, s.requester.GetUserId())

	if err != nil {
		return shared.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return usermodel.ErrUserHasBeenBlock
	}

	taskPerDay, err := s.tStore.Count(ctx, map[string]interface{}{
		"created_at": shared.FormatDateString(time.Now()),
		"user_id":    s.requester.GetUserId(),
	})

	if *taskPerDay > 5 {
		return taskmodel.ErrYouAreLimitedReach
	}

	if err != nil {
		return shared.ErrCannotGetEntity(taskmodel.EntityName, err)
	}

	taskDataModel := taskmodel.Task{
		Content: data.Content,
		UserId:  s.requester.GetUserId(),
	}
	if err = s.tStore.Create(ctx, &taskDataModel); err != nil {
		return shared.ErrCannotCreateEntity(taskmodel.EntityName, err)
	}
	return nil
}
