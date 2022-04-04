package service

import (
	"github.com/ansidev/togo/domain/task"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/task/dto"
	taskMock "github.com/ansidev/togo/task/mock"
	"github.com/ansidev/togo/test"
	userMock "github.com/ansidev/togo/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func TestTaskService(t *testing.T) {
	suite.Run(t, new(TaskServiceTestSuite))
}

type TaskServiceTestSuite struct {
	test.ServiceTestSuite
	mockUserRepository *userMock.MockIUserRepository
	mockTaskRepository *taskMock.MockITaskRepository
}

func (s *TaskServiceTestSuite) SetupSuite() {
	s.ServiceTestSuite.SetupSuite()
	s.mockUserRepository = userMock.NewMockIUserRepository(s.Ctrl)
	s.mockTaskRepository = taskMock.NewMockITaskRepository(s.Ctrl)
}

func (s *TaskServiceTestSuite) TestCreate_WithTotalToDayTaskReachedLimit_ShouldReturnReachedLimitDailyTaskError() {
	maxDailyTask := 5
	userModel := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2",
		MaxDailyTask: maxDailyTask,
		CreatedAt:    time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 1, 23, 56, 0, time.UTC),
	}

	createTaskRequest := dto.CreateTaskRequest{Title: "Task 1"}

	mockTotalTodayTask := int64(maxDailyTask + rand.Intn(10))

	s.mockUserRepository.
		EXPECT().
		FindFirstByID(userModel.ID).
		Return(userModel, nil)

	s.mockTaskRepository.
		EXPECT().
		GetTotalTasksByUserAndDate(userModel, gomock.AssignableToTypeOf(time.Time{})).
		Return(mockTotalTodayTask, nil)

	taskService := NewTaskService(s.mockUserRepository, s.mockTaskRepository)
	_, err := taskService.Create(createTaskRequest, userModel.ID)

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrReachedLimitDailyTask, errs.Message(err))
	require.Equal(s.T(), errs.ErrCodeReachedLimitDailyTask, errs.ErrorCode(err))
}

func (s *TaskServiceTestSuite) TestCreate_WhenUserDoesNotExists_ShouldReturnUsernameNotFoundError() {
	maxDailyTask := 10
	userModel := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2",
		MaxDailyTask: maxDailyTask,
		CreatedAt:    time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 1, 23, 56, 0, time.UTC),
	}

	s.mockUserRepository.
		EXPECT().
		FindFirstByID(int64(1)).
		Return(user.User{}, errors.Wrap(errs.ErrRecordNotFound, gorm.ErrRecordNotFound.Error()))

	createTaskRequest := dto.CreateTaskRequest{Title: "Task 1"}

	taskService := NewTaskService(s.mockUserRepository, s.mockTaskRepository)
	_, err := taskService.Create(createTaskRequest, userModel.ID)

	require.Error(s.T(), err)
	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrUsernameNotFound, errs.Message(err))
	require.Equal(s.T(), errs.ErrCodeUsernameNotFound, errs.ErrorCode(err))
}

func (s *TaskServiceTestSuite) TestCreate_WithTotalToDayTaskLessThanMaxDailyTask_ShouldReturnCreatedTask() {
	maxDailyTask := 10
	userModel := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2",
		MaxDailyTask: maxDailyTask,
		CreatedAt:    time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 1, 23, 56, 0, time.UTC),
	}

	mockTotalTodayTask := int64(rand.Intn(maxDailyTask))

	s.mockUserRepository.
		EXPECT().
		FindFirstByID(userModel.ID).
		Return(userModel, nil)

	s.mockTaskRepository.
		EXPECT().
		GetTotalTasksByUserAndDate(userModel, gomock.AssignableToTypeOf(time.Time{})).
		Return(mockTotalTodayTask, nil)

	createTaskRequest := dto.CreateTaskRequest{Title: "Task 1"}

	mockCreatedTask := task.Task{
		Title:     createTaskRequest.Title,
		UserID:    userModel.ID,
		CreatedAt: time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
		UpdatedAt: time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}
	mockCreatedTask.ID = int64(rand.Intn(100))
	s.mockTaskRepository.
		EXPECT().
		Create(gomock.AssignableToTypeOf(task.Task{}), userModel).
		Return(mockCreatedTask, nil)

	taskService := NewTaskService(s.mockUserRepository, s.mockTaskRepository)
	createdTask, err := taskService.Create(createTaskRequest, userModel.ID)

	require.NoError(s.T(), err)
	require.Equal(s.T(), mockCreatedTask.ID, createdTask.ID)
	require.Equal(s.T(), mockCreatedTask.Title, createdTask.Title)
	require.Equal(s.T(), mockCreatedTask.UserID, createdTask.OwnerID)

	_, err1 := time.Parse(time.RFC3339, createdTask.CreatedAt)
	_, err2 := time.Parse(time.RFC3339, createdTask.UpdatedAt)
	require.NoError(s.T(), err1)
	require.NoError(s.T(), err2)
}
