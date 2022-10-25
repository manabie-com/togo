package repository_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/backend/entity"
	mockRepo "github.com/manabie-com/backend/mocks/repository"
	taskservice "github.com/manabie-com/backend/services/task"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	task := entity.Task{
		ID:          uuid.NewV4().String(),
		UserID:      uuid.NewV4().String(),
		Content:     "Task1",
		Status:      "pendding",
		CreatedDate: utils.GetToday(),
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := mockRepo.NewMockI_Repository(ctl)
	mockRepo.EXPECT().CreateTask(&task).Return(nil)

	service := taskservice.Service{Repo: mockRepo}

	result, err := service.CreateTask(&task)

	assert.Nil(t, nil, err)
	assert.Equal(t, task.ID, result.ID)
	assert.Equal(t, task.UserID, result.UserID)
	assert.Equal(t, task.Content, result.Content)
	assert.Equal(t, task.Status, result.Status)
	assert.Equal(t, task.CreatedDate, result.CreatedDate)
}
