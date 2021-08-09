package domain

import (
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain/model"
	"github.com/manabie-com/togo/repo"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func Test_todoService_GetTaskAtDate(t1 *testing.T) {
	userRepo := new(repo.UserRepositoryMock)
	taskRepo := new(repo.TaskRepositoryMock)
	contextMock := new(context.ContextMock)
	now := time.Now()
	nowRounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	type fields struct {
		jwtKey       string
		tokenTimeout int
		userRepo     repo.UserRepository
		taskRepo     repo.TaskRepository
	}
	type args struct {
		ctx        context.Context
		createDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				jwtKey:       "wqGyEBBfPK9w3Lxw",
				tokenTimeout: 60,
				userRepo:     userRepo,
				taskRepo:     taskRepo,
			},
			args: args{
				ctx:        contextMock,
				createDate: nowRounded,
			},
			want: []*model.Task{
				{
					ID:          "c8545a7f-81a4-4774-af41-c8fbe64afaa4",
					Content:     "some task",
					UserID:      "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
					CreatedDate: nowRounded,
				},
			},
		},
	}
	contextMock.On("GetUserId").Return("6e135b1d-1e54-46f3-99f8-5b78df9cb857")
	taskRepo.On("FindTaskByUserIdAndDate", contextMock, "6e135b1d-1e54-46f3-99f8-5b78df9cb857", nowRounded).Return([]*model.Task{
		{
			ID:          "c8545a7f-81a4-4774-af41-c8fbe64afaa4",
			Content:     "some task",
			UserID:      "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
			CreatedDate: nowRounded,
		},
	}, nil)
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todoService{
				jwtKey:       tt.fields.jwtKey,
				tokenTimeout: tt.fields.tokenTimeout,
				userRepo:     tt.fields.userRepo,
				taskRepo:     tt.fields.taskRepo,
			}
			got, err := t.GetTaskAtDate(tt.args.ctx, tt.args.createDate)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTaskAtDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTaskAtDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_CreateTask(t1 *testing.T) {
	userRepo := new(repo.UserRepositoryMock)
	taskRepo := new(repo.TaskRepositoryMock)
	contextMock := new(context.ContextMock)
	now := time.Now()
	nowRounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	type fields struct {
		jwtKey       string
		tokenTimeout int
		userRepo     repo.UserRepository
		taskRepo     repo.TaskRepository
	}
	type args struct {
		ctx     context.Context
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				jwtKey:       "wqGyEBBfPK9w3Lxw",
				tokenTimeout: 60,
				userRepo:     userRepo,
				taskRepo:     taskRepo,
			},
			args: args{
				ctx:     contextMock,
				content: "abcd",
			},
			want: &model.Task{
				CreatedDate: nowRounded,
			},
		},
	}

	contextMock.On("GetUserId").Return("6e135b1d-1e54-46f3-99f8-5b78df9cb857")
	taskRepo.On("FindTaskByUserIdAndDate", contextMock, "6e135b1d-1e54-46f3-99f8-5b78df9cb857", nowRounded).Return([]*model.Task{
		{
			ID:          "c8545a7f-81a4-4774-af41-c8fbe64afaa4",
			Content:     "some task",
			UserID:      "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
			CreatedDate: nowRounded,
		},
	}, nil)
	taskRepo.On("Insert", contextMock, mock.Anything).Return(nil)
	userRepo.On("GetUserById", contextMock, "6e135b1d-1e54-46f3-99f8-5b78df9cb857").Return(&model.User{
		Id:       "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
		Username: "abcd",
		Password: "efgh",
		MaxTodo:  2,
	}, nil)

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todoService{
				jwtKey:       tt.fields.jwtKey,
				tokenTimeout: tt.fields.tokenTimeout,
				userRepo:     tt.fields.userRepo,
				taskRepo:     tt.fields.taskRepo,
			}
			got, err := t.CreateTask(tt.args.ctx, tt.args.content)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && tt.want.CreatedDate != got.CreatedDate {
				t1.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_CreateTask_Fail(t1 *testing.T) {
	userRepo := new(repo.UserRepositoryMock)
	taskRepo := new(repo.TaskRepositoryMock)
	contextMock := new(context.ContextMock)
	now := time.Now()
	nowRounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	type fields struct {
		jwtKey       string
		tokenTimeout int
		userRepo     repo.UserRepository
		taskRepo     repo.TaskRepository
	}
	type args struct {
		ctx     context.Context
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "fail because limit task per day",
			fields: fields{
				jwtKey:       "wqGyEBBfPK9w3Lxw",
				tokenTimeout: 60,
				userRepo:     userRepo,
				taskRepo:     taskRepo,
			},
			args: args{
				ctx:     contextMock,
				content: "abcd",
			},
			wantErr: true,
		},
	}

	contextMock.On("GetUserId").Return("6e135b1d-1e54-46f3-99f8-5b78df9cb857")
	taskRepo.On("FindTaskByUserIdAndDate", contextMock, "6e135b1d-1e54-46f3-99f8-5b78df9cb857", nowRounded).Return([]*model.Task{
		{
			ID:          "c8545a7f-81a4-4774-af41-c8fbe64afaa4",
			Content:     "some task",
			UserID:      "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
			CreatedDate: nowRounded,
		},
	}, nil)
	taskRepo.On("Insert", contextMock, mock.Anything).Return(nil)
	userRepo.On("GetUserById", contextMock, "6e135b1d-1e54-46f3-99f8-5b78df9cb857").Return(&model.User{
		Id:       "6e135b1d-1e54-46f3-99f8-5b78df9cb857",
		Username: "abcd",
		Password: "efgh",
		MaxTodo:  1,
	}, nil)

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todoService{
				jwtKey:       tt.fields.jwtKey,
				tokenTimeout: tt.fields.tokenTimeout,
				userRepo:     tt.fields.userRepo,
				taskRepo:     tt.fields.taskRepo,
			}
			got, err := t.CreateTask(tt.args.ctx, tt.args.content)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && tt.want.CreatedDate != got.CreatedDate {
				t1.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
