package handler

import (
	"errors"
	"testing"
	"time"

	"togo-internal-service/internal/model"
	"togo-internal-service/mock/mock_storage"
	v1 "togo-internal-service/pkg/api/v1"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_CreateTask(t *testing.T) {

	tests := []struct {
		handler Handler
		name    string
		req     *v1.CreateTaskRequest
		want    *v1.Task
		wantErr bool
	}{
		{
			name: "validate error title empty",
			req: &v1.CreateTaskRequest{
				UserId:  "user-01",
				Content: "the content",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: nil,
				Config:  &Config{},
			},
		},
		{
			name: "validate error user_id empty",
			req: &v1.CreateTaskRequest{
				Title:   "the title",
				Content: "the content",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: nil,
				Config:  &Config{},
			},
		},
		{
			name: "validate error created_time nil",
			req: &v1.CreateTaskRequest{
				UserId:      "user-01",
				Title:       "the title",
				Content:     "the content",
				CreatedTime: nil,
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: nil,
				Config:  &Config{},
			},
		},
		{
			name: "storage create error",
			req: &v1.CreateTaskRequest{
				UserId:      "user-01",
				Title:       "the title",
				Content:     "the content",
				CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					CreateTaskFunc: func(ctx context.Context, task *model.Task) (*model.Task, error) {
						return nil, errors.New("some error happended")
					},
				},
				Config: &Config{},
			},
		},
		{
			name: "happy case #0",
			req: &v1.CreateTaskRequest{
				UserId:      "user-01",
				Title:       "the title",
				Content:     "the content",
				CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
			},
			want: &v1.Task{
				UserId:      "user-01",
				Title:       "the title",
				Content:     "the content",
				CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
			},
			wantErr: false,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					CreateTaskFunc: func(ctx context.Context, task *model.Task) (*model.Task, error) {
						return &model.Task{
							UserID:      task.UserID,
							Title:       task.Title,
							Content:     task.Content,
							CreatedTime: task.CreatedTime,
						}, nil
					},
				},
				Config: &Config{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			resp, err := test.handler.CreateTask(ctx, test.req)
			if (err != nil) != test.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !assert.True(t, proto.Equal(resp, test.want)) {
				t.Errorf("CreateTask() resp = %v, want %v", resp, test.want)
				return
			}
		})
	}
}
