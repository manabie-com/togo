package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func Test_NewRedis(t *testing.T) {
	s := NewRedis(RedisClient)

	assert.IsType(t, &Redis{}, s)
}

func Test_SetExpire(t *testing.T) {
	key := "1"
	value := 1
	expire := time.Duration(1)
	testCases := []struct{
		name string
		setupMock func(mr redismock.ClientMock) redismock.ClientMock
		assertErr func(err error)
	}{
		{
			name: "when there is internal error, then should return error",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectSetEX(key, value, expire).SetErr(errors.New("internal redis error"))
				return mr
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "internal redis error")
			},
		},
		{
			name: "when there is no error, then should return nil",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectSetEX(key, value, expire).SetVal("OK")
				return mr
			},
			assertErr: func(err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			mock = test.setupMock(mock)
			redis := NewRedis(db)
			err := redis.SetExpire(key, value, expire)
			test.assertErr(err)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_GetInt(t *testing.T) {
	key := "1"
	testCases := []struct{
		name string
		cacheValue string
		setupMock func(mr redismock.ClientMock) redismock.ClientMock
		assertErr func(err error)
		expectedLimit int
	}{
		{
			name: "when there is not error, then should return limit number",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectGet(key).SetVal("2")
				return mr
			},
			assertErr: func(err error) {
				assert.Nil(t, err)
			},
			expectedLimit: 2,
		},
		{
			name: "when there is invalid integer value, then should invalid integer error",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectGet(key).SetVal("a")
				return mr
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "invalid integer")
			},
			expectedLimit: 0,
		},
		{
			name: "when there is redis nil value, then should return zero value with nil error",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectGet(key).RedisNil()
				return mr
			},
			assertErr: func(err error) {
				assert.Nil(t, err)
			},
			expectedLimit: 0,
		},
		{
			name: "when there is internal redis error, then should return zero value with non-nil error",
			setupMock: func(mr redismock.ClientMock) redismock.ClientMock {
				mr.ExpectGet(key).SetErr(errors.New("internal redis error"))
				return mr
			},
			assertErr: func(err error) {
				assert.ErrorContains(t, err, "internal redis error")
			},
			expectedLimit: 0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			mock = test.setupMock(mock)
			redis := NewRedis(db)
			limit, err := redis.GetInt(key)
			assert.Equal(t, test.expectedLimit, limit)
			test.assertErr(err)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}