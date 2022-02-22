package redisrepo

import (
	"context"
	"fmt"
	"testing"
	"time"
	"togo/internal/domain"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func Test_taskLimitRepository_Increase_FirstInTheDay(t *testing.T) {
	prefix := "task_limit"
	user := &domain.User{
		ID:          1,
		TasksPerDay: 1,
	}
	ct := time.Now().UTC()
	exp := ct.AddDate(0, 1, -ct.Day())
	key := fmt.Sprintf("%s:%v:%v", prefix, user.ID, ct.Day())
	db, mock := redismock.NewClientMock()
	r := taskLimitRepository{
		rdb:    db,
		prefix: "task_limit",
	}
	mock.ExpectGet(key).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectIncr(key).SetVal(1)
	mock.ExpectExpireAt(key, exp).SetVal(true)
	mock.ExpectTxPipelineExec()
	currentTasksSubmitted, err := r.Increase(context.Background(), user.ID, user.TasksPerDay)
	assert.Equal(t, 1, currentTasksSubmitted)
	assert.NoError(t, err)
}

func Test_taskLimitRepository_Increase_LimitExceed(t *testing.T) {
	prefix := "task_limit"
	user := &domain.User{
		ID:          1,
		TasksPerDay: 1,
	}
	ct := time.Now().UTC()
	key := fmt.Sprintf("%s:%v:%v", prefix, user.ID, ct.Day())
	db, mock := redismock.NewClientMock()
	r := taskLimitRepository{
		rdb:    db,
		prefix: "task_limit",
	}
	mock.ExpectGet(key).SetVal("1")
	currentTasksSubmitted, err := r.Increase(context.Background(), user.ID, user.TasksPerDay)
	assert.Equal(t, 1, currentTasksSubmitted)
	assert.ErrorIs(t, err, domain.ErrTaskLimitExceed)
}
