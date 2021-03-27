package usecases

import (
	"context"
	"github.com/manabie-com/togo/domains"
	"github.com/stretchr/testify/mock"
	"net/http"
)

// Prepare mock for auth
type AuthMock struct {
	mock.Mock
}

func (a *AuthMock) CreateToken(userId int64) (string, error) {
	args := a.Called(userId)
	return args.String(0), args.Error(1)
}

func (a *AuthMock) ValidateToken(req *http.Request) (int64, error) {
	args := a.Called(req)
	return args.Get(0).(int64), args.Error(1)
}

// Prepare db mock for repositories
type DBMock struct {
	mock.Mock
}

// mocks for Task repository
func (m *DBMock) GetCountCreatedTaskTodayByUser(ctx context.Context, userId int64) (int64, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(int64), args.Error(1)
}

func (m *DBMock) CreateTask(ctx context.Context, taskInput *domains.Task) (*domains.Task, error) {
	args := m.Called(ctx, taskInput)
	var task *domains.Task
	if args.Get(0) != nil {
		task = args.Get(0).(*domains.Task)
	}
	return task, args.Error(1)
}

func (m *DBMock) GetTasks(ctx context.Context, request *domains.TaskRequest) ([]*domains.Task, error) {
	args := m.Called(ctx, request)
	var tasks []*domains.Task
	if args.Get(0) != nil {
		tasks = args.Get(0).([]*domains.Task)
	}
	return tasks, args.Error(1)
}

func (m *DBMock) GetTaskById(ctx context.Context, request *domains.TaskByIdRequest) (*domains.Task, error) {
	args := m.Called(ctx, request)
	var task *domains.Task
	if args.Get(0) != nil {
		task = args.Get(0).(*domains.Task)
	}
	return task, args.Error(1)
}

// =================================
// mocks for User repository
func (m *DBMock) VerifyUser(ctx context.Context, request *domains.LoginRequest) (*domains.User, error) {
	args := m.Called(ctx, request)
	var user *domains.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domains.User)
	}
	return user, args.Error(1)
}

func (m *DBMock) GetUserById(ctx context.Context, userId int64) (*domains.User, error) {
	args := m.Called(ctx, userId)
	var user *domains.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domains.User)
	}
	return user, args.Error(1)
}
