package services

import (
	"testing"

	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/repositories"
)

func TestAddTaskToUser(t *testing.T) {
	userMockRepo := repositories.NewUserMockRepository()
	userTaskMockRepo := repositories.NewUserTaskMockRepository()
	userTaskSrv := NewUserTaskService(userTaskMockRepo, userMockRepo)

	testCases := []testCase{
		{
			name: "Task should be created successfully.",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 1",
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      "2022-02-17",
			},
			expected: map[string]interface{}{
				"info": map[string]interface{}{
					"title":       "Unit Test Task 1",
					"description": "My unit test task one",
					"user_name":   "Test User 1",
				},
			},
		},
		{
			name: "validator.ValidationErrors must be returned when the given UserName is empty",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    " ", // Empty UserName field
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      "2022-02-17",
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "validator.ValidationErrors must be returned when the given Title is empty",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 1",
				Title:       " ", // Empty Title field
				Description: "My unit test task one",
				InsDay:      "2022-02-17",
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "validator.ValidationErrors must be returned when the given Description is empty",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 1",
				Title:       "Unit Test Task 1",
				Description: " ", // Empty Description field
				InsDay:      "2022-02-17",
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "validator.ValidationErrors must be returned when the given InsDay is empty",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 1",
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      " ", // Empty InsDay field
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "validator.ValidationErrors must be returned when the given InsDay is not in the format of YYYY-MM-DD",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 1",
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      "Wrong Format", // Wrong InsDay format
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "*apierrors.UserDoesNotExistsError must be returned when the given UserName does not exists.",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Unit Test User 1", // This user does not exists in UserTaskMockRepository
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      "2022-02-17",
			},
			hasError: true,
			errType:  UserDoesNotExistsError,
		},
		{
			name: "*apierrors.MaxTasksReachedError must be returned when tasks was added in a user that already reached the maximum number of tasks.",
			createTaskDTO: dto.CreateTaskDTO{
				UserName:    "Test User 3", // This user already has already reached the maximum tasks per day (2)
				Title:       "Unit Test Task 1",
				Description: "My unit test task one",
				InsDay:      "2022-02-17",
			},
			hasError: true,
			errType:  MaxTasksReachedError,
		},
	}

	for _, tc := range testCases {
		res, err := userTaskSrv.AddTaskToUser(tc.createTaskDTO)
		assertTestCase(t, tc, res, err)
	}

}

func TestGetTasksOfUser(t *testing.T) {
	userMockRepo := repositories.NewUserMockRepository()
	userTaskMockRepo := repositories.NewUserTaskMockRepository()
	userTaskSrv := NewUserTaskService(userTaskMockRepo, userMockRepo)

	testCases := []testCase{
		{
			name: "Should return the tasks of user",
			getTaskOfUserDTO: dto.GetTaskOfUserDTO{
				UserName: "Test User 1",
				InsDay:   "2022-02-17",
			},
			expected: map[string]interface{}{
				"user_task": map[string]interface{}{
					"user_id":   "620e6b6e20bdcb887326931a",
					"user_name": "Test User 1",
					"max_tasks": 3,
					"ins_day":   "2022-02-17",
					"tasks": []map[string]interface{}{
						{
							"title":       "User 1 Task 1 02-17",
							"description": "User One Task One",
						},
						{
							"title":       "User 1 Task 2 02-17",
							"description": "User One Task Two",
						},
					},
				},
			},
		},
		{
			name: "Should return an empty tasks field when the user does not have a tasks on the given InsDay",
			getTaskOfUserDTO: dto.GetTaskOfUserDTO{
				UserName: "Test User 1",
				InsDay:   "2022-02-18", // Test user 1 does not have tasks on "2022-02-18"
			},
			expected: map[string]interface{}{
				"user_task": map[string]interface{}{
					"user_id":   "620e6b6e20bdcb887326931a",
					"user_name": "Test User 1",
					"max_tasks": 3,
					"ins_day":   "2022-02-18",
					"tasks":     []map[string]interface{}{},
				},
			},
		},
		{
			name: "validator.ValidationErrors must be returned when the given UserName is empty",
			getTaskOfUserDTO: dto.GetTaskOfUserDTO{
				UserName: " ", // Empty UserName field
				InsDay:   "2022-02-17",
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "*apierrors.UserDoesNotExistsError must be returned when the given UserName does not exists.",
			getTaskOfUserDTO: dto.GetTaskOfUserDTO{
				UserName: "Unit Test User 1", // This user does not exists
				InsDay:   "2022-02-17",
			},
			hasError: true,
			errType:  UserDoesNotExistsError,
		},
	}

	for _, tc := range testCases {
		res, err := userTaskSrv.GetTasksOfUser(tc.getTaskOfUserDTO)
		assertTestCase(t, tc, res, err)
	}
}
