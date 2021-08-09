package repo

import (
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain/model"
	"github.com/stretchr/testify/mock"
	"time"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (t *TaskRepositoryMock) Insert(ctx context.Context, task *model.Task) error {
	args := t.Called(ctx, task)
	return args.Error(0)
}

func (t *TaskRepositoryMock) FindTaskByUserIdAndDate(ctx context.Context, userId string, createdDate time.Time) ([]*model.Task, error) {
	args := t.Called(ctx, userId, createdDate)
	return args.Get(0).([]*model.Task), args.Error(1)
}
