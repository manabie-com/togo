package repositories

import "github.com/kier1021/togo/api/models"

type UserTaskMockRepository struct {
	userTask []models.UserTask
}

func NewUserTaskMockRepository() *UserTaskMockRepository {

	userTask := []models.UserTask{
		{
			UserID: "620e6b6e20bdcb887326931a",
			User: models.User{
				ID:       "620e6b6e20bdcb887326931a",
				UserName: "Test User 1",
				MaxTasks: 3,
			},
			Tasks: []models.Task{
				{
					Title:       "User 1 Task 1 02-17",
					Description: "User One Task One",
				},
				{
					Title:       "User 1 Task 2 02-17",
					Description: "User One Task Two",
				},
			},
			InsDay: "2022-02-17",
		},
		{
			UserID: "620e6b79657f405b5f79b32d",
			User: models.User{
				ID:       "620e6b79657f405b5f79b32d",
				UserName: "Test User 2",
				MaxTasks: 4,
			},
			Tasks: []models.Task{
				{
					Title:       "User 2 Task 1 02-17",
					Description: "User Two Task One",
				},
				{
					Title:       "User 2 Task 2 02-17",
					Description: "User Two Task Two",
				},
				{
					Title:       "User 2 Task Three 02-17",
					Description: "User Two Task Three",
				},
			},
			InsDay: "2022-02-17",
		},
		{
			UserID: "620e6b7e64b5c80f08aaddcd",
			User: models.User{
				ID:       "620e6b7e64b5c80f08aaddcd",
				UserName: "Test User 3",
				MaxTasks: 2,
			},
			Tasks: []models.Task{
				{
					Title:       "User 3 Task 1 02-17",
					Description: "User Three Task One",
				},
				{
					Title:       "User 3 Task 2 02-17",
					Description: "User Three Task Two",
				},
			},
			InsDay: "2022-02-17",
		},
	}

	return &UserTaskMockRepository{
		userTask: userTask,
	}
}

func (repo *UserTaskMockRepository) AddTaskToUser(user models.User, userTask models.Task, insDay string) error {
	return nil
}

func (repo *UserTaskMockRepository) GetUserTask(filter map[string]interface{}) (userTask *models.UserTask, err error) {
	users := repo.filterUserTasks(filter)

	if len(users) != 0 {
		userTask = &users[0]
	}

	return userTask, nil
}

func (repo *UserTaskMockRepository) filterUserTasks(filter map[string]interface{}) (users []models.UserTask) {
	for _, u := range repo.userTask {

		isEqual := false

		if userID, ok := filter["user_id"]; ok {
			isEqual = u.UserID == userID
		}

		if userName, ok := filter["user_name"]; ok {
			isEqual = u.UserName == userName
		}

		if insDay, ok := filter["ins_day"]; ok {
			isEqual = u.InsDay == insDay
		}

		if isEqual {
			users = append(users, u)
		}

	}

	return users
}
