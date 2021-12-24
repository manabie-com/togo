package handler

import (
	"testing"
	"time"

	"togo-internal-service/internal/model"
	"togo-internal-service/internal/storage"
	"togo-internal-service/mock/mock_storage"
	v1 "togo-internal-service/pkg/api/v1"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_GetTask(t *testing.T) {
	tests := []struct {
		handler Handler
		name    string
		req     *v1.GetTaskRequest
		want    *v1.Task
		wantErr bool
	}{
		{
			name: "get db error",
			req: &v1.GetTaskRequest{
				Id: "",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					GetTaskFunc: func(ctx context.Context, ID string) (*model.Task, error) {
						return nil, storage.ErrTaskNotFound
					},
				},
			},
		},
		{
			name: "happy case #0",
			req: &v1.GetTaskRequest{
				Id: "123",
			},
			want: &v1.Task{
				Id:          "123",
				Title:       "the title",
				Content:     "the content",
				UserId:      "user-01",
				CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
			},
			wantErr: false,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					GetTaskFunc: func(ctx context.Context, ID string) (*model.Task, error) {
						return &model.Task{
							ID:          ID,
							Title:       "the title",
							Content:     "the content",
							UserID:      "user-01",
							CreatedTime: time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
						}, nil
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			resp, err := test.handler.GetTask(ctx, test.req)
			if (err != nil) != test.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !assert.True(t, proto.Equal(resp, test.want)) {
				t.Errorf("GetTask() resp = %v, want %v", resp, test.want)
				return
			}
		})
	}
}
