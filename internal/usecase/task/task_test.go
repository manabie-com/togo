package task_test

import (
	"errors"
	"testing"

	"github.com/manabie-com/togo/internal/mocks"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase/task"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

var (
	task1 = storages.Task{
		Content: "",
		UserID:  "secondUser",
	}
)

func TestRetrieveTasks(t *testing.T) {
	defer goleak.VerifyNone(t)

	testCases := []struct {
		context     string
		buildStubs  func(taskStore *mocks.MockTaskStorageInterface)
		user        storages.User
		expectedErr error
		expected    []*storages.Task
	}{
		{
			context: "success",
			buildStubs: func(taskStore *mocks.MockTaskStorageInterface) {
				taskStore.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return([]*storages.Task{&task1, &task1}, nil)
			},
			expectedErr: nil,
			expected:    []*storages.Task{&task1, &task1},
		},
		{
			context: "cannot find user",
			buildStubs: func(taskStore *mocks.MockTaskStorageInterface) {
				taskStore.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return([]*storages.Task{}, errors.New("cannot retrieve tasks"))
			},
			expectedErr: errors.New("cannot retrieve tasks"),
			expected:    []*storages.Task{},
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			taskStore := mocks.NewMockTaskStorageInterface(controller)
			c.buildStubs(taskStore)

			taskUsecase := task.NewTaskUsecase(taskStore)

			tasks, err := taskUsecase.RetrieveTasks(c.user.ID, c.user.Password)
			assert.Equal(t, c.expectedErr, err)
			assert.Equal(t, len(c.expected), len(tasks))
		})
	}
}
