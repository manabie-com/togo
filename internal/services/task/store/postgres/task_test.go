package postgres

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"togo/internal/services/task/domain"
	"togo/internal/services/task/store"
	"togo/pkg/random"
)

func createTask(t *testing.T) *domain.Task {
	user := createUser(t)
	id, _ := uuid.NewUUID()
	title := random.RandomQuote()
	description := random.RandomQuote()
	dueDate := time.Now()
	task := domain.NewTask(id, user.ID, title, description, dueDate)
	err := taskRepo.Save(task)
	assert.NoError(t, err)
	return task
}

func TestTaskRepository_Save(t *testing.T) {
	t.Run("should save task", func(t *testing.T) {
		createTask(t)
	})
	t.Run("should error if title is empty", func(t *testing.T) {
		user := createUser(t)
		id, _ := uuid.NewUUID()
		task := domain.NewTask(id, user.ID, "", "", time.Now())
		err := taskRepo.Save(task)
		assert.Error(t, err)
	})
	t.Run("should error if user is not exists", func(t *testing.T) {
		id, _ := uuid.NewUUID()
		title := random.RandomQuote()
		description := random.RandomQuote()
		dueDate := time.Now()
		task := domain.NewTask(id, uuid.Nil, title, description, dueDate)
		err := taskRepo.Save(task)
		assert.Error(t, err)
	})
}

func TestTaskRepository_Count(t *testing.T) {
	t.Run("should count tasks with user id", func(t *testing.T) {
		numberOfTasks := random.RandomInt(1, 10)
		user := createUser(t)
		now := time.Now()
		for i := 0; i < numberOfTasks; i++ {
			task := domain.NewTask(uuid.New(), user.ID, random.RandomQuote(), random.RandomQuote(), now)
			err := taskRepo.Save(task)
			assert.NoError(t, err)
		}
		count, err := taskRepo.Count(store.CountTasksRequest{
			UserID: &user.ID,
		})
		assert.NoError(t, err)
		assert.Equal(t, numberOfTasks, count)
	})
	t.Run("should count tasks with user id and day", func(t *testing.T) {
		task := createTask(t)
		numberOfTasks := random.RandomInt(1, 10)
		tomorrow := task.DueDate.Add(24 * time.Hour)
		for i := 0; i < numberOfTasks; i++ {
			tomorrowTask := domain.NewTask(uuid.New(), task.UserID, random.RandomQuote(), random.RandomQuote(), tomorrow)
			err := taskRepo.Save(tomorrowTask)
			assert.NoError(t, err)
		}
		count, err := taskRepo.Count(store.CountTasksRequest{
			UserID: &task.UserID,
			Day:    &tomorrow,
		})
		assert.NoError(t, err)
		assert.Equal(t, numberOfTasks, count)
	})
}
