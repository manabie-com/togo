package storage

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/lock/etcd"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages/psql"
	"github.com/stretchr/testify/assert"
)

func setupStorage(t *testing.T) *psql.Storage {
	var conf struct {
		Psql psql.Config
		Etcd etcd.Config
	}
	panicIfErr(envconfig.Process("storage", &conf))
	s, err := psql.NewStorage(conf.Psql)
	panicIfErr(err)
	etcd, err := etcd.NewLock(conf.Etcd)
	panicIfErr(err)
	s.WithLock(etcd)
	return s
}

var (
	fakeUsers = []domain.User{
		{
			ID:             "1",
			Password:       "1",
			MaxTasksPerDay: 1,
		},
		{
			ID:             "2",
			Password:       "2",
			MaxTasksPerDay: 2,
		},
		{
			ID:             "3",
			Password:       "3",
			MaxTasksPerDay: 3,
		},
	}
)

func TestPostgresql(t *testing.T) {
	s := setupStorage(t)
	defer func() {
		err := s.CleanupDB()
		panicIfErr(err)
	}()
	t.Run("Add user", func(t *testing.T) {
		for _, u := range fakeUsers {
			err := s.CreateUser(u)
			assert.NoError(t, err)
			newu, err := s.FindUserByID(u.ID)
			assert.NoError(t, err)
			assert.Equal(t, u.ID, newu.ID)
			assert.Equal(t, u.Password, newu.Password)
			assert.Equal(t, u.MaxTasksPerDay, newu.MaxTasksPerDay)
		}
	})
	t.Run("Add duplicate user", func(t *testing.T) {
		for _, u := range fakeUsers {
			err := s.CreateUser(u)
			assert.Error(t, err)
		}
	})
	t.Run("Get user tasks per day", func(t *testing.T) {
		for _, u := range fakeUsers {
			perday, err := s.GetUserTasksPerDay(u.ID)
			assert.NoError(t, err)
			assert.Equal(t, u.MaxTasksPerDay, perday)
		}
	})
	t.Run("Concurrently adding tasks", func(t *testing.T) {
		u := makeSuperTaskUser()
		err := s.CreateUser(u)
		assert.NoError(t, err)
		wg := sync.WaitGroup{}
		totalRequest := u.MaxTasksPerDay * 2
		wg.Add(totalRequest)
		todate := time.Now().Format(domain.DateFormat)

		var totalLimitReached = int32(0)

		for i := 0; i < totalRequest; i++ {
			go func() {
				defer wg.Done()
				err := s.AddTaskWithLimitPerDay(makeFakeTask(u.ID, todate), u.MaxTasksPerDay)
				if err != nil {
					// ignore these cases
					if errors.Is(err, psql.ErrTooManySerializableConflict) {
						return
					}
					//count total tasks denied
					if errors.Is(err, domain.TaskLimitReached) {
						atomic.AddInt32(&totalLimitReached, 1)
						return
					}
					assert.FailNowf(t, "Unexpected error on adding tasks", "error: %s", err)
				}
			}()
		}
		wg.Wait()
		ts, err := s.GetTasksByUserIDAndDate(u.ID, todate)
		assert.NoError(t, err)
		if totalLimitReached > 0 {
			assert.Equal(t, u.MaxTasksPerDay, len(ts), "Task created (actual) equal limitation (expected)")
		} else {
			assert.LessOrEqual(t, len(ts), u.MaxTasksPerDay)
		}

	})
}

func makeSuperTaskUser() domain.User {
	return domain.User{
		ID:             uuid.New().String(),
		Password:       "password",
		MaxTasksPerDay: 10,
	}
}

func makeFakeTask(userID, todate string) domain.Task {
	return domain.Task{
		ID:          uuid.New().String(),
		UserID:      userID,
		Content:     "hello",
		CreatedDate: todate,
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
