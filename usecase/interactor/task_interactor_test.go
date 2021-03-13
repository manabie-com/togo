package interactor

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/mocks"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/request"
	"github.com/valonekowd/togo/usecase/response"
)

func Test_basicTaskInteractor_Fetch(t *testing.T) {
	testCases := []struct {
		name          string
		returnTasks   []*entity.Task
		repoError     error
		returnResp    *response.GetTasks
		expectedError error
	}{
		{
			name: "Number of tasks greater than zero",
			returnTasks: []*entity.Task{
				{
					ID:        "this-is-id",
					Content:   "this-is-content",
					UserID:    "this-is-user-id",
					CreatedAt: time.Now().UTC(),
				},
			},
			repoError: nil,
			returnResp: &response.GetTasks{
				Data: []*response.TaskData{
					{
						ID:          "this-is-id",
						Content:     "this-is-content",
						UserID:      "this-is-user-id",
						CreatedDate: "this-is-created-date",
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Number of tasks is zero",
			returnTasks:   nil,
			repoError:     nil,
			returnResp:    &response.GetTasks{},
			expectedError: nil,
		},
		{
			name:          "Repository return error",
			returnTasks:   nil,
			repoError:     errors.New("whoops"),
			returnResp:    nil,
			expectedError: fmt.Errorf("fetching tasks: %w", errors.New("whoops")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			taskRepo := mocks.TaskRepositoryMock{}
			taskRepo.On("List", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(tc.returnTasks, tc.repoError)

			taskFormatter := mocks.TaskFormatterMock{}
			taskFormatter.On("Fetch", mock.MatchedBy(func(ctx context.Context) bool { return true }), tc.returnTasks).Return(tc.returnResp)

			i := basicTaskInteractor{
				ds:        interfaces.DataSource{Task: &taskRepo},
				presenter: &taskFormatter,
			}
			got, err := i.Fetch(context.Background(), request.GetTasks{})
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			if got != nil {
				assert.Equal(t, len(tc.returnTasks), len(got.Data))
			}
		})
	}
}

func Test_TaskInteractor_Create(t *testing.T) {
	fakeTask := &entity.Task{
		ID:        "this-is-id",
		Content:   "this-is-content",
		UserID:    "this-is-user-id",
		CreatedAt: time.Now().UTC(),
	}

	testCases := []struct {
		name          string
		task          *entity.Task
		repoError     error
		returnResp    *response.CreateTask
		expectedError error
	}{
		{
			name:      "Create task successful",
			task:      fakeTask,
			repoError: nil,
			returnResp: &response.CreateTask{
				Data: &response.TaskData{
					ID:          "this-is-id",
					Content:     "this-is-content",
					UserID:      "this-is-user-id",
					CreatedDate: "this-is-created-date",
				},
			},
			expectedError: nil,
		},
		{
			name:          "Create task and repository return error",
			task:          fakeTask,
			repoError:     errors.New("whoops"),
			returnResp:    nil,
			expectedError: fmt.Errorf("creating task: %w", errors.New("whoops")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			taskRepo := mocks.TaskRepositoryMock{}
			taskRepo.On("Add", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.MatchedBy(func(t *entity.Task) bool { return t.Content == tc.task.Content })).Return(tc.repoError)

			taskFormatter := mocks.TaskFormatterMock{}
			taskFormatter.On("Create", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.MatchedBy(func(t *entity.Task) bool { return t.Content == tc.task.Content })).Return(tc.returnResp)

			i := basicTaskInteractor{
				ds:        interfaces.DataSource{Task: &taskRepo},
				presenter: &taskFormatter,
			}
			got, err := i.Create(context.Background(), request.CreateTask{Content: tc.task.Content})
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			if got != nil {
				assert.Equal(t, tc.returnResp, got)
			}
		})
	}
}
