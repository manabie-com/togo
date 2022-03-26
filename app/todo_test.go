package app_test

import (
	"context"
	"testing"

	mock "github.com/golang/mock/gomock"
	"github.com/laghodessa/togo/app"
	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/test/todofixture"
	"github.com/laghodessa/togo/test/todomock"
	"github.com/stretchr/testify/require"
)

func TestTodoUsecase_AddTask(t *testing.T) {
	ctrl := mock.NewController(t)
	defer ctrl.Finish()

	userRepo := todomock.NewMockUserRepo(ctrl)
	taskRepo := todomock.NewMockTaskRepo(ctrl)

	userRepo.EXPECT().
		GetUser(mock.Any(), mock.Any()).
		Return(todofixture.NewUser(), nil)

	taskRepo.EXPECT().
		AddTask(mock.Any(), mock.Any(), mock.Any()).
		Return(nil)

	taskRepo.EXPECT().
		CountInTimeRangeByUserID(mock.Any(), mock.Any(), mock.Any(), mock.Any()).
		Return(0, nil)

	uc := &app.TodoUsecase{
		TaskRepo: taskRepo,
		UserRepo: userRepo,
	}

	req := app.AddTask{
		Task: app.Task{
			UserID:  domain.NewID(),
			Message: "stop coding",
		},
		TimeZone: "Asia/Ho_Chi_Minh",
	}
	_, err := uc.AddTask(context.Background(), req)
	require.NoError(t, err)
}
