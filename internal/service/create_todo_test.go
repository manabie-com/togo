package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chi07/todo/internal/model"
	"github.com/chi07/todo/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks2 "github.com/chi07/todo/tests/mocks"
)

func TestCreateTaskService_Create(t *testing.T) {
	ctx := context.Background()
	taskID := uuid.New()
	limitMockRepo := new(mocks2.LimitationRepo)
	taskMockRepo := new(mocks2.CreateTaskRepo)

	tests := []struct {
		name    string
		task    *model.Task
		wantErr error
		mock    func()
		want    uuid.UUID
	}{
		{
			name: "error: when get user limit post failed",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			want: uuid.Nil,
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(nil, errors.New("internal error"))
			},
			wantErr: errors.New("internal error"),
		},
		{
			name: "when a user have no limitation",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(nil, nil)
			},
			want:    uuid.Nil,
			wantErr: errors.New("no record found"),
		},
		{
			name: "when a count user tasks was failed",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(&model.Limitation{
					ID:        0,
					UserID:    uuid.New().String(),
					LimitTask: 5,
				}, nil)
				taskMockRepo.On("CountUserTasks", ctx, mock.Anything).Once().Return(int64(0), errors.New("internal err"))
			},
			want:    uuid.Nil,
			wantErr: errors.New("internal err"),
		},
		{
			name: "when user got max daily tasks",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(&model.Limitation{
					ID:        1,
					UserID:    uuid.New().String(),
					LimitTask: 5,
				}, nil)
				taskMockRepo.On("CountUserTasks", ctx, mock.Anything).Once().Return(int64(5), nil)
			},
			want:    uuid.Nil,
			wantErr: errors.New("permission denied, you reach maximum posts a day"),
		},
		{
			name: "when create task got failed",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(&model.Limitation{
					ID:        1,
					UserID:    uuid.New().String(),
					LimitTask: 5,
				}, nil)
				taskMockRepo.On("CountUserTasks", ctx, mock.Anything).Once().Return(int64(2), nil)
				taskMockRepo.On("Create", ctx, mock.Anything).Once().Return(uuid.Nil, errors.New(" cannot create task"))
			},
			want:    uuid.Nil,
			wantErr: errors.New("cannot create task:  cannot create task"),
		},
		{
			name: "success",
			task: &model.Task{
				Title:    "Test",
				Status:   "TODO",
				Priority: "URGENT",
				UserID:   uuid.New(),
			},
			mock: func() {
				limitMockRepo.On("GetByUserID", ctx, mock.Anything).Once().Return(&model.Limitation{
					ID:        1,
					UserID:    uuid.New().String(),
					LimitTask: 5,
				}, nil)
				taskMockRepo.On("CountUserTasks", ctx, mock.Anything).Once().Return(int64(1), nil)
				taskMockRepo.On("Create", ctx, mock.Anything).Once().Return(taskID, nil)
			},
			want: taskID,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			taskService := service.NewCreateTaskService(taskMockRepo, limitMockRepo)
			actual, err := taskService.CreateTask(ctx, tc.task)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("taskService.CreateTask() got %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}
