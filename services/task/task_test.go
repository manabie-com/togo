package task

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/backend/entity"
	mockService "github.com/manabie-com/backend/mocks/taskservice"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func CallService(serv I_TaskService, task *entity.Task) (*entity.Task, *utils.ErrorRest) {
	newTask, err := serv.CreateTask(task)

	return newTask, err
}
func CallServiceGetAll(serv I_TaskService, tasks []entity.Task) ([]entity.Task, *utils.ErrorRest) {
	newTask, err := serv.GetTaskAll()

	return newTask, err
}

func CreateTask() entity.Task {
	return entity.Task{
		ID:          uuid.NewV4().String(),
		UserID:      uuid.NewV4().String(),
		Content:     "Task1",
		Status:      "pendding",
		CreatedDate: utils.GetToday(),
	}
}

func TestTaskService(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	task := CreateTask()

	taskServiceMock := mockService.NewMockI_TaskService(ctl)
	taskServiceMock.EXPECT().CreateTask(&task).Return(&task, nil)

	result, err := CallService(taskServiceMock, &task)

	assert.Nil(t, nil, err)
	assert.Equal(t, task.ID, result.ID)
	assert.Equal(t, task.UserID, result.UserID)
	assert.Equal(t, task.Content, result.Content)
	assert.Equal(t, task.Status, result.Status)
	assert.Equal(t, task.CreatedDate, result.CreatedDate)
}

func TestServiceGetTaskAll(t *testing.T) {
	arrFiveTasks := make([]entity.Task, 0)

	for i := 0; i < 5; i++ {
		newTask := CreateTask()
		arrFiveTasks = append(arrFiveTasks, newTask)
	}

	testCases := []struct {
		name       string
		arrTask    []entity.Task
		buildStubs func(mock *mockService.MockI_TaskService, tasks []entity.Task)
		check      func(t *testing.T, tasks []entity.Task, error interface{})
	}{
		{
			name:    "Should be empty array",
			arrTask: make([]entity.Task, 0),
			buildStubs: func(mock *mockService.MockI_TaskService, tasks []entity.Task) {

				mock.EXPECT().GetTaskAll().Return(tasks, nil)
			},
			check: func(t *testing.T, tasks []entity.Task, error interface{}) {
				assert.Nil(t, error)
				assert.Equal(t, 0, len(tasks))
			},
		},
		{
			name:    "Should be five ",
			arrTask: arrFiveTasks,
			buildStubs: func(mock *mockService.MockI_TaskService, tasks []entity.Task) {

				mock.EXPECT().GetTaskAll().Return(tasks, nil)
			},
			check: func(t *testing.T, tasks []entity.Task, error interface{}) {
				assert.Nil(t, error)
				assert.Equal(t, 5, len(tasks))
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		ctl := gomock.NewController(t)
		defer ctl.Finish()

		taskServiceMock := mockService.NewMockI_TaskService(ctl)
		testCase.buildStubs(taskServiceMock, testCase.arrTask)

		result, err := CallServiceGetAll(taskServiceMock, testCase.arrTask)
		testCase.check(t, result, err)
	}
}
