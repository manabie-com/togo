package storages

import (
	"context"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/utils"
	"time"
)

//TODO Replace Mock service with :https://github.com/golang/mock
type MockUserRepository struct {
}

func (mockRepo *MockUserRepository) ValidateUser(ctx context.Context, userID, pwd string) bool {
	if userID == "test" && pwd == "test" {
		return true
	}

	return false
}

type MockTaskRepository struct {
	tasks []*entities.Task
}

func (mockRepo *MockTaskRepository) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entities.Task, error) {
	var result []*entities.Task
	for _, task := range mockRepo.tasks {
		if task.UserID == userID && task.CreatedDate == createdDate {
			result = append(result, task)
		}
	}
	return result, nil
}

func (mockRepo *MockTaskRepository) AddTask(ctx context.Context, t *entities.Task) error {
	mockRepo.tasks = append(mockRepo.tasks, t)
	return nil
}

func (mockRepo *MockTaskRepository) CountTaskPerDayByUserID(ctx context.Context, userID string) (uint, error) {
	count := 0
	for _, task := range mockRepo.tasks {
		if task.UserID == userID && task.CreatedDate == utils.FormatTimeToString(time.Now()) {
			count++
		}
	}
	return uint(count), nil
}
