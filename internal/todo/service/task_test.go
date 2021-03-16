package service

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/manabie-com/togo/internal/todo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskService_ListTaskForUser(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		userID  int
		dateStr string
	}
	tests := []struct {
		name     string
		taskRepo d.TaskRepository
		args     args
		want     []*d.Task
		wantErr  bool
	}{
		{
			"System Error",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("GetTasksForUser", mock.Anything, 1, mock.Anything).
					Return(nil, errors.New("System Error"))
				return repo
			}(),
			args{userID: 1, dateStr: "2021-03-17"},
			nil,
			true,
		},
		{
			"Empty date",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("GetTasksForUser", mock.Anything, 2, time.Now().Format("2006-01-02")).
					Return([]*d.Task{
						{ID: 1},
						{ID: 2},
					}, nil)
				return repo
			}(),
			args{userID: 2, dateStr: ""},
			[]*d.Task{
				{ID: 1},
				{ID: 2},
			},
			false,
		},
		{
			"Invalid date",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("GetTasksForUser", mock.Anything, 2, time.Now().Format("2006-01-02")).
					Return([]*d.Task{
						{ID: 1},
						{ID: 2},
					}, nil)
				return repo
			}(),
			args{userID: 2, dateStr: "abcdef"},
			[]*d.Task{
				{ID: 1},
				{ID: 2},
			},
			false,
		},
		{
			"Valid result",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("GetTasksForUser", mock.Anything, 2, "2021-03-19").
					Return([]*d.Task{
						{ID: 1},
						{ID: 2},
					}, nil)
				return repo
			}(),
			args{userID: 2, dateStr: "2021-03-19"},
			[]*d.Task{
				{ID: 1},
				{ID: 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				TaskRepo: tt.taskRepo,
			}

			ctx := context.Background()
			got, err := s.ListTaskForUser(ctx, tt.args.userID, tt.args.dateStr)
			mockRepo, _ := s.TaskRepo.(*mocks.TaskRepository)

			mockRepo.AssertExpectations(t)
			assert.Equal(tt.want, got)
			assert.True((err != nil) == tt.wantErr)
		})
	}
}

func TestTaskService_CreateTaskForUser(t *testing.T) {
	assert := assert.New(t)
	os.Setenv("MAX_TASKS_DAILY", "3")
	type args struct {
		userID int
		param  d.TaskCreateParam
	}
	tests := []struct {
		name     string
		taskRepo d.TaskRepository
		args     args
		want     *d.Task
		err      error
	}{
		{
			"System Error create",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("CreateTaskForUser", mock.Anything, 1, d.TaskCreateParam{Content: "test"}).
					Return(nil, errors.New("System Error"))
				return repo
			}(),
			args{userID: 1, param: d.TaskCreateParam{Content: "test"}},
			nil,
			errors.New("System Error"),
		},
		{
			"Max tasks daily reached",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("CreateTaskForUser", mock.Anything, 1, d.TaskCreateParam{Content: "test"}).
					Return(nil, nil)
				return repo
			}(),
			args{userID: 1, param: d.TaskCreateParam{Content: "test"}},
			nil,
			d.ErrTaskLimitReached,
		},
		{
			"Tasks created",
			func() d.TaskRepository {
				repo := &mocks.TaskRepository{}
				repo.On("CreateTaskForUser", mock.Anything, 1, d.TaskCreateParam{Content: "test"}).
					Return(&d.Task{UserID: 1, Content: "test"}, nil)
				return repo
			}(),
			args{userID: 1, param: d.TaskCreateParam{Content: "test"}},
			&d.Task{UserID: 1, Content: "test"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				TaskRepo: tt.taskRepo,
			}

			ctx := context.Background()
			got, err := s.CreateTaskForUser(ctx, tt.args.userID, tt.args.param)
			mockRepo, _ := s.TaskRepo.(*mocks.TaskRepository)

			mockRepo.AssertExpectations(t)
			assert.Equal(tt.want, got)
			assert.Equal(tt.err, err)
		})
	}
}
