package user_test

import (
	"errors"
	"testing"

	"github.com/manabie-com/togo/internal/mocks"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase/user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

var (
	user1 = storages.User{
		ID:       "firstUser",
		Password: "example",
		MaxTodo:  5,
	}

	task = storages.Task{
		Content: "",
		UserID:  "secondUser",
	}

	user2 = storages.User{
		ID:       "secondUser",
		Password: "password",
		MaxTodo:  2,
		Tasks:    []storages.Task{task, task},
	}
)

func TestCreateTask(t *testing.T) {
	defer goleak.VerifyNone(t)

	testCases := []struct {
		context     string
		buildStubs  func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface)
		task        storages.Task
		expectedErr error
	}{
		{
			context: "success",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUsersTasks(gomock.Any(), gomock.Any()).Return(user1, nil)
				taskStore.EXPECT().CreateTask(gomock.Any()).Return(nil)
			},
			expectedErr: nil,
		},
		{
			context: "get user task error",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUsersTasks(gomock.Any(), gomock.Any()).Return(storages.User{}, errors.New("Cannot find user"))
			},
			expectedErr: errors.New("Cannot find user"),
		},
		{
			context: "over limit today",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUsersTasks(gomock.Any(), gomock.Any()).Return(user2, nil)
			},
			expectedErr: errors.New("over limit per day"),
		},
		{
			context: "create task fail",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUsersTasks(gomock.Any(), gomock.Any()).Return(user1, nil)
				taskStore.EXPECT().CreateTask(gomock.Any()).Return(errors.New("create task fail"))
			},
			expectedErr: errors.New("create task fail"),
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			userStore := mocks.NewMockUserStorageInterface(controller)
			taskStore := mocks.NewMockTaskStorageInterface(controller)
			c.buildStubs(userStore, taskStore)

			userUsecase := user.NewUserUsecase(userStore, taskStore)

			err := userUsecase.CreateTask(&c.task)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}

func TestValidateUser(t *testing.T) {
	defer goleak.VerifyNone(t)

	testCases := []struct {
		context     string
		buildStubs  func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface)
		user        storages.User
		expectedErr error
	}{
		{
			context: "success",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedErr: nil,
		},

		{
			context: "cannot find user",
			buildStubs: func(userStore *mocks.MockUserStorageInterface, taskStore *mocks.MockTaskStorageInterface) {
				userStore.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(errors.New("Cannot find user"))
			},
			expectedErr: errors.New("Cannot find user"),
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			userStore := mocks.NewMockUserStorageInterface(controller)
			taskStore := mocks.NewMockTaskStorageInterface(controller)
			c.buildStubs(userStore, taskStore)

			userUsecase := user.NewUserUsecase(userStore, taskStore)

			err := userUsecase.ValidateUser(c.user.ID, c.user.Password)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
