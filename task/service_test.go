package task

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func Test_NewService(t *testing.T) {
	s := NewService(new(mockRepository), new(mockCacheService))

	assert.IsType(t, &Service{}, s)
}

var mockTimeNow = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

func Test_RecordTaskSuccess(t *testing.T)  {
	repo := new(mockRepository)
	cacheService := new(mockCacheService)
	loc, _ := time.LoadLocation("Local")
	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, loc)
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 99, loc)
	duration := midnight.Sub(now)
	userId := "1"
	monkey.Patch(time.Now, func() time.Time {
		return now
	})

	repo.On("GetUserConfig", userId).Return(&UserConfig{
		UserId: "1",
		Limit: 3,
	}, nil)

	repo.On("InsertUserTask", userId,"todo", time.Now().In(loc)).Return(nil)

	cacheService.On("GetInt", userId).Return(2, nil)
	cacheService.On("SetExpire", userId, 3, duration).Return(nil)
	recordService := NewService(repo, cacheService)

	err := recordService.RecordTask(userId, "todo")
	assert.Nil(t, err)
}