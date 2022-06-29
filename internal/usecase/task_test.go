package usecase

import (
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/pkg"
	mockdb "lntvan166/togo/pkg/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var task1 = &e.Task{
	ID:          1,
	Name:        "task",
	Description: "task",
	CreatedAt:   pkg.GetCurrentTime(),
	Completed:   false,
	UserID:      1,
}

var task2 = &e.Task{
	ID:          2,
	Name:        "task",
	Description: "task",
	CreatedAt:   pkg.GetCurrentTime(),
	Completed:   false,
	UserID:      1,
}

var user = &e.User{
	ID:       1,
	Username: "user",
	Password: "user",
	Plan:     "free",
	MaxTodo:  10,
}

var tasks = &[]e.Task{
	*task1,
	*task2,
}

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)
	userRepo.EXPECT().GetUserByID(user.ID).Return(user, nil)
	taskRepo.EXPECT().GetNumberOfTaskTodayByUserID(user.ID).Return(0, nil).AnyTimes()
	taskRepo.EXPECT().CreateTask(task1).Return(nil)

	id, err := taskUsecase.CreateTask(task1, user.Username)
	assert.NoError(t, err)
	assert.Equal(t, 0, id) // number of task today = 3
}

func TestGetAllTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	taskRepo.EXPECT().GetAllTask().Return(tasks, nil)

	returnTasks, err := taskUsecase.GetAllTask()
	assert.NoError(t, err)
	assert.Equal(t, tasks, returnTasks)
}

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	taskRepo.EXPECT().GetTaskByID(task1.ID).Return(task1, nil).AnyTimes()
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)

	returnTask, err := taskUsecase.GetTaskByID(task1.ID, user.Username)
	assert.NoError(t, err)
	assert.Equal(t, task1, returnTask)
}

func TestGetTasksByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)
	taskRepo.EXPECT().GetTasksByUserID(user.ID).Return(tasks, nil)

	returnTasks, err := taskUsecase.GetTasksByUsername(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, tasks, returnTasks)
}

func TestGetUserIDByTaskID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	taskRepo.EXPECT().GetTaskByID(task1.ID).Return(task1, nil)

	returnUserID, err := taskUsecase.GetUserIDByTaskID(task1.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, returnUserID)
}

func TestCheckLimitTaskToday(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	userRepo.EXPECT().GetUserByID(user.ID).Return(user, nil)
	taskRepo.EXPECT().GetNumberOfTaskTodayByUserID(user.ID).Return(3, nil)

	limit, err := taskUsecase.CheckLimitTaskToday(task1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, false, limit)
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil).AnyTimes()
	taskRepo.EXPECT().GetTaskByID(task1.ID).Return(task1, nil).AnyTimes()
	taskRepo.EXPECT().UpdateTask(task1).Return(nil)

	err := taskUsecase.UpdateTask(task1.ID, user.Username, *task1)
	assert.NoError(t, err)
}

func TestCompleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	taskRepo.EXPECT().GetTaskByID(task1.ID).Return(task1, nil)
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)
	taskRepo.EXPECT().CompleteTask(task1.ID).Return(nil)

	err := taskUsecase.CompleteTask(task1.ID, user.Username)
	assert.NoError(t, err)
}

func TestCheckAccessPermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)

	err := taskUsecase.CheckAccessPermission(user.Username, task1.ID)
	assert.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	userRepo := mockdb.NewMockUserRepository(ctrl)

	taskUsecase := NewTaskUsecase(taskRepo, userRepo)

	// before test
	taskRepo.EXPECT().GetTaskByID(task1.ID).Return(task1, nil)
	userRepo.EXPECT().GetUserIDByUsername(user.Username).Return(1, nil)
	taskRepo.EXPECT().DeleteTask(task1.ID).Return(nil)

	err := taskUsecase.DeleteTask(task1.ID, user.Username)
	assert.NoError(t, err)
}
