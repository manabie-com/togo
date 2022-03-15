package biz_test

import (
	"errors"
	"testing"

	"github.com/HoangMV/todo/src/biz"
	"github.com/HoangMV/todo/src/models/entity"
	"github.com/HoangMV/todo/src/models/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DaoMock struct {
	mock.Mock
}

func (m *DaoMock) CountUserTodoInCurrentDay(userID int) (int, error) {
	ret := m.Called(userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *DaoMock) GetMaxUserTodoOneDay(userID int) (int, error) {
	ret := m.Called(userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *DaoMock) InsertTodo(obj *entity.Todo) error {
	ret := m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Todo) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *DaoMock) InsertUser(obj *entity.User) error                       { return nil }
func (m *DaoMock) InsertUserMaxTodo(obj *entity.UserTodoConfig) error      { return nil }
func (m *DaoMock) UpdateTodo(obj *entity.Todo) error                       { return nil }
func (m *DaoMock) GetUserByUsername(username string) (*entity.User, error) { return nil, nil }
func (m *DaoMock) SelectTodosByUserID(userID int, size, index int) ([]entity.Todo, error) {
	return nil, nil
}
func (m *DaoMock) GetTokenInCache(username string) int      { return 0 }
func (m *DaoMock) SetTokenToCache(token string, userID int) {}

type Require struct {
	cur int
	max int
}

func TestCreateTodo(t *testing.T) {

	testCases := []struct {
		name     string
		in       *request.CreateTodoReq
		require  Require
		expected error
	}{
		{
			"CreateTodo failed",
			&request.CreateTodoReq{
				UserID:  1,
				Content: "aaaaaaaa",
			},
			Require{
				cur: 2,
				max: 2,
			},
			errors.New("your todo count has reached its maximum"),
		},
		{
			"CreateTodo success",
			&request.CreateTodoReq{
				UserID:  1,
				Content: "aaaaaaaa",
			},
			Require{
				cur: 1,
				max: 2,
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			daoMock := new(DaoMock)
			daoMock.On("CountUserTodoInCurrentDay", tc.in.UserID).Return(tc.require.cur, nil)
			daoMock.On("GetMaxUserTodoOneDay", tc.in.UserID).Return(tc.require.max, nil)
			daoMock.On("InsertTodo", &entity.Todo{UserID: tc.in.UserID, Content: tc.in.Content}).Return(nil)

			biz := biz.NewWithDao(daoMock)
			err := biz.CreateTodo(tc.in)

			if tc.expected != nil {
				if assert.Error(t, err) {
					assert.Equal(t, tc.expected, err)
				}
			}
		})
	}

}
