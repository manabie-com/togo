package record

import (
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_NewService(t *testing.T) {
	s := NewService(new(mockRepository), new(mockCacheService))

	assert.IsType(t, &Service{}, s)
}

func Test_RecordTaskSuccess(t *testing.T)  {
	loc, _ := time.LoadLocation("Local")
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, loc)
	duration := midnight.Sub(now)
	monkey.Patch(time.Now, func() time.Time {
		return now
	})

	testUserId := "1"
	testConfigLimit := 3
	testCacheLimit := 2
	testCases := []struct{
		name string
		setupMock func(mr *mockRepository, mc *mockCacheService)
		assertErr func(err error)
	}{
		{
			name: "when no errors happen, then the record should be successful",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, nil)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(testCacheLimit, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "when get user config error happens, then the record should be failed",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, errors.New("internal mongo db error"))
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(testCacheLimit, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "get user config error")
			},
		},
		{
			name: "when user does not exist in db, then the record should be failed with user not exist error",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, mongo.ErrNoDocuments)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(testCacheLimit, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "user id does not exist")
			},
		},
		{
			name: "when there is an error with get redis key, then the record should be failed with get redis error",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, nil)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(testCacheLimit, errors.New("internal redis get failed"))
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "get redis error")
			},
		},
		{
			name: "when the limit reached, then the record should be failed with user record record reached limit error",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, nil)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(3, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "user record record reached limit")
			},
		},
		{
			name: "when insert user's record gets error, then the record should be failed with insert user's record error",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, nil)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(errors.New("internal mongo insert error"))

				mc.On("GetInt", testUserId).Return(2, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(nil)
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "insert user's record error")
			},
		},
		{
			name: "when set expire gets error, then the record should be failed with set redis expire error",
			setupMock: func(mr *mockRepository, mc *mockCacheService) {
				mr.On("GetUserConfig", testUserId).Return(&UserConfig{
					UserId: testUserId,
					Limit: testConfigLimit,
				}, nil)
				mr.On("InsertUserTask", testUserId,"todo", time.Now().In(loc)).Return(nil)

				mc.On("GetInt", testUserId).Return(2, nil)
				mc.On("SetExpire", testUserId, testCacheLimit + 1, duration).Return(errors.New("internal redis set expire error"))
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "set redis expire error")
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			repo := new(mockRepository)
			cacheService := new(mockCacheService)
			test.setupMock(repo, cacheService)
			recordService := NewService(repo, cacheService)
			err := recordService.RecordTask(testUserId, "todo")
			test.assertErr(err)
		})
	}
} 