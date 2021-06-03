package tasks

import (
	"context"
	"testing"

	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
)

// Mock test service, used for unit testing
type MockTaskService struct {
	Testing                      *testing.T
	CreateNewResponse            error
	GetTasksCreatedOnResponse    []*models.Task
	TaskCountByUserResponseCount int64
	TaskCountByUserResponseError error

	ShouldCreateNewCalled         bool
	ShouldGetTasksCreatedOnCalled bool
	ShouldTaskCountByUserCalled   bool
}

func (m MockTaskService) CreateNew(context.Context, *models.Task) error {
	if !m.ShouldCreateNewCalled && m.Testing != nil {
		assert.Fail(m.Testing, "CreateNew function is called, expected not to be called")
	}
	return m.CreateNewResponse
}

func (m MockTaskService) GetTasksCreatedOn(context.Context, string) []*models.Task {
	if !m.ShouldGetTasksCreatedOnCalled && m.Testing != nil {
		assert.Fail(m.Testing, "GetTasksCreatedOn function is called, expected not to be called")
	}
	return m.GetTasksCreatedOnResponse
}

func (m MockTaskService) TaskCountByUser(context.Context, string) (int64, error) {
	if !m.ShouldTaskCountByUserCalled && m.Testing != nil {
		assert.Fail(m.Testing, "TaskCountByUser function is called, expected not to be called")
	}
	return m.TaskCountByUserResponseCount, m.TaskCountByUserResponseError
}
