package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/mocks"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func taskTestSetupMockTaskStore(t *testing.T, s *mocks.MockTaskStore) *mocks.MockTaskStore {
	local := make(map[string][]domain.Task)
	s.EXPECT().AddTaskWithLimitPerDay(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(t domain.Task, limit int) error {
		tasks, exist := local[t.UserID]
		if !exist {
			if defaultTaskPerDay > 0 {
				local[t.UserID] = []domain.Task{t}
			} else {
				return fmt.Errorf("mock error")
			}
		}
		if len(tasks) >= defaultTaskPerDay {
			return domain.TaskLimitReached{}
		}
		tasks = append(tasks, t)
		local[t.UserID] = tasks
		return nil
	})
	s.EXPECT().GetTasksByUserID(gomock.Any()).AnyTimes().DoAndReturn(func(uid string) ([]domain.Task, error) {
		ts, exist := local[uid]
		if !exist {
			return nil, nil
		}
		return ts, nil
	})
	return s
}

func taskTestSetupMockUserStore(t *testing.T, s *mocks.MockUserStore) *mocks.MockUserStore {
	local := make(map[string]domain.User)
	s.EXPECT().CreateUser(gomock.Any()).AnyTimes().DoAndReturn(func(u domain.User) error {
		_, exist := local[u.ID]
		if exist {
			return fmt.Errorf("mock error")
		}
		local[u.ID] = u
		return nil
	})
	s.EXPECT().GetUserTasksPerDay(gomock.Any()).AnyTimes().DoAndReturn(func(id string) (int, error) {
		u, exist := local[id]
		if !exist {
			return 0, domain.UserNotFound(id)
		}
		return u.MaxTasksPerDay, nil
	})
	return s
}

var (
	defaultTaskPerDay = 5
	templateTask      = domain.Task{
		Content: "hello",
		UserID:  "admin",
	}
)

func TestTaskUC(t *testing.T) {
	c := gomock.NewController(t)
	mockUserStore := mocks.NewMockUserStore(c)
	mockUserStore = taskTestSetupMockUserStore(t, mockUserStore)
	mockTaskStore := mocks.NewMockTaskStore(c)
	mockTaskStore = taskTestSetupMockTaskStore(t, mockTaskStore)

	assert.NoError(t, mockUserStore.CreateUser(domain.User{
		ID:             "admin",
		Password:       "admin",
		MaxTasksPerDay: defaultTaskPerDay,
	}))
	uc := usecase.NewTaskUseCase(mockTaskStore, mockUserStore)
	for i := 0; i < defaultTaskPerDay; i++ {
		task := templateTask
		task.ID = uuid.New().String()
		task.CreatedDate = time.Now().Format(domain.DateFormat)
		err := uc.AddTask(task)
		assert.NoError(t, err)
	}
	tasks, err := uc.GetTasksByUserID("admin")
	assert.NoError(t, err)
	assert.Equal(t, defaultTaskPerDay, len(tasks))
	for _, item := range tasks {
		assert.Equal(t, templateTask.Content, item.Content)
		assert.Equal(t, templateTask.UserID, item.UserID)
	}
	task := templateTask
	task.ID = uuid.New().String()
	task.CreatedDate = time.Now().Format(domain.DateFormat)
	err = uc.AddTask(task)
	assert.ErrorIs(t, err, domain.TaskLimitReached{})

}
