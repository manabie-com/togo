package taskservice

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"todo/configs"
	"todo/internal/entities"
	"todo/internal/repository"
	"todo/internal/repository/mock"
	"todo/pkg/helper"
	"reflect"
	"testing"
)

func Test_TaskServiceCreateTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type fields struct {
		taskRepo repository.TaskRepository
		cfg      *configs.Config
	}
	type args struct {
		req *entities.Task
		ctx context.Context
	}
	var (
		ctx     = context.Background()
		ctxUser = context.WithValue(context.Background(), "user", &jwt.Token{
			Claims: &helper.JwtClaims{
				UserID:   1,
				UserName: "daiphuc",
			},
		})
		err  = errors.New("got error")
		req  = &entities.Task{ID: 1, Name: "Task 1", Content: "Content 1", UserID: 1}
		resp = &entities.Task{ID: 1, Name: "Task 1", Content: "Content 1", UserID: 1}
	)

	tests := []struct {
		name    string
		args    args
		fields  fields
		errResp error
		want    *entities.Task
	}{
		{
			name: "get user from context failed",
			args: args{
				req: req,
				ctx: ctx,
			},
			fields:  fields{},
			errResp: helper.ErrUserNotFound,
			want:    nil,
		},
		{
			name: "create task failed",
			args: args{
				req: req,
				ctx: ctxUser,
			},
			fields: fields{
				taskRepo: func() *mock.MockTaskRepository {
					mockRepo := mock.NewMockTaskRepository(mockCtrl)
					mockRepo.EXPECT().Create(ctxUser, req).Return(resp, nil)
					return mockRepo
				}(),
			},
			errResp: nil,
			want:    resp,
		},
		{
			name: "create task success",
			args: args{
				req: req,
				ctx: ctxUser,
			},
			fields: fields{
				taskRepo: func() *mock.MockTaskRepository {
					mockRepo := mock.NewMockTaskRepository(mockCtrl)
					mockRepo.EXPECT().Create(ctxUser, req).Return(nil, err)
					return mockRepo
				}(),
			},
			errResp: err,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTaskService(tt.fields.taskRepo)
			got, errOut := ts.CreateTask(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.errResp, errOut)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
