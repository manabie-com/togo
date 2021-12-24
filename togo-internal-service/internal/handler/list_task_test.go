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
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_ListTask(t *testing.T) {
	tests := []struct {
		handler Handler
		name    string
		req     *v1.ListTaskRequest
		want    *v1.ListTaskResponse
		wantErr bool
	}{
		{
			name: "validate error user_id empty",
			req: &v1.ListTaskRequest{
				UserId:    "",
				Date:      &date.Date{},
				PageSize:  1,
				PageToken: "",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Config: &Config{
					MaxListTaskPageSize: 10,
				},
			},
		},
		{
			name: "validate error page_size invalid #0",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      &date.Date{},
				PageSize:  0,
				PageToken: "",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Config: &Config{
					MaxListTaskPageSize: 10,
				},
			},
		},
		{
			name: "validate error page_size invalid #1",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      &date.Date{},
				PageSize:  -1,
				PageToken: "",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Config: &Config{
					MaxListTaskPageSize: 10,
				},
			},
		},
		{
			name: "validate error page_token invalid",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      &date.Date{},
				PageSize:  1,
				PageToken: "this is invalid token",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Config: &Config{
					MaxListTaskPageSize: 10,
				},
			},
		},
		{
			name: "error get db",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      nil,
				PageSize:  1,
				PageToken: "",
			},
			want:    nil,
			wantErr: true,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					ListTaskFunc: func(ctx context.Context, userID string, date time.Time, limit int, offset int) ([]*model.Task, error) {
						return nil, storage.ErrInternal
					},
				},
				Config: &Config{
					MaxListTaskPageSize: 10,
				},
			},
		},
		{
			name: "happy case #0",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      &date.Date{},
				PageSize:  2,
				PageToken: "",
			},
			want: &v1.ListTaskResponse{
				Tasks: []*v1.Task{
					{
						Id:          "1",
						Title:       "title 1",
						Content:     "content 1",
						UserId:      "user-01",
						CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
					},
					{
						Id:          "2",
						Title:       "title 2",
						Content:     "content 2",
						UserId:      "user-01",
						CreatedTime: timestamppb.New(time.Date(2, 2, 2, 2, 2, 2, 2, time.UTC)),
					},
				},
				NextPageToken: model.PagingToToken(&model.Paging{
					Offset: 2,
				}),
			},
			wantErr: false,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					ListTaskFunc: func(ctx context.Context, userID string, date time.Time, limit int, offset int) ([]*model.Task, error) {
						return []*model.Task{
							{
								ID:          "1",
								Title:       "title 1",
								Content:     "content 1",
								UserID:      userID,
								CreatedTime: time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
							},
							{
								ID:          "2",
								Title:       "title 2",
								Content:     "content 2",
								UserID:      userID,
								CreatedTime: time.Date(2, 2, 2, 2, 2, 2, 2, time.UTC),
							},
						}, nil
					},
				},
				Config: &Config{
					MaxListTaskPageSize: 1,
				},
			},
		},
		{
			name: "happy case #1",
			req: &v1.ListTaskRequest{
				UserId:    "user-01",
				Date:      &date.Date{},
				PageSize:  2,
				PageToken: "",
			},
			want: &v1.ListTaskResponse{
				Tasks: []*v1.Task{
					{
						Id:          "1",
						Title:       "title 1",
						Content:     "content 1",
						UserId:      "user-01",
						CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
					},
				},
				NextPageToken: "",
			},
			wantErr: false,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					ListTaskFunc: func(ctx context.Context, userID string, date time.Time, limit int, offset int) ([]*model.Task, error) {
						return []*model.Task{
							{
								ID:          "1",
								Title:       "title 1",
								Content:     "content 1",
								UserID:      userID,
								CreatedTime: time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
							},
						}, nil
					},
				},
				Config: &Config{
					MaxListTaskPageSize: 3,
				},
			},
		},
		{
			name: "happy case #3",
			req: &v1.ListTaskRequest{
				UserId:   "user-01",
				Date:     &date.Date{},
				PageSize: 1,
				PageToken: model.PagingToToken(&model.Paging{
					Offset: 1,
				}),
			},
			want: &v1.ListTaskResponse{
				Tasks: []*v1.Task{
					{
						Id:          "1",
						Title:       "title 1",
						Content:     "content 1",
						UserId:      "user-01",
						CreatedTime: timestamppb.New(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)),
					},
				},
				NextPageToken: model.PagingToToken(&model.Paging{
					Offset: 2,
				}),
			},
			wantErr: false,
			handler: Handler{
				Storage: &mock_storage.StorageMock{
					ListTaskFunc: func(ctx context.Context, userID string, date time.Time, limit int, offset int) ([]*model.Task, error) {
						return []*model.Task{
							{
								ID:          "1",
								Title:       "title 1",
								Content:     "content 1",
								UserID:      userID,
								CreatedTime: time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
							},
						}, nil
					},
				},
				Config: &Config{
					MaxListTaskPageSize: 3,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			resp, err := test.handler.ListTask(ctx, test.req)
			if (err != nil) != test.wantErr {
				t.Errorf("ListTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !assert.True(t, proto.Equal(resp, test.want)) {
				t.Errorf("ListTask()\nresp = %v\nwant = %v", resp, test.want)
				return
			}
		})
	}
}
