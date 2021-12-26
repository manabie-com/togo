package handler

import (
	"context"
	"errors"
	"testing"

	"togo-public-api/internal/auth"
	"togo-public-api/internal/service/togo_internal_v1"
	"togo-public-api/internal/service/togo_user_session_v1"
	"togo-public-api/mock/mock_togo_internal_v1"
	"togo-public-api/mock/mock_togo_user_session_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func Test_GetTask(t *testing.T) {
	tests := []struct {
		name     string
		req      *v1.GetTaskRequest
		token    string
		want     *v1.Task
		wantErr  bool
		internal togo_internal_v1.TogoInternalServiceClient
		session  togo_user_session_v1.TogoUserSessionServiceClient
	}{
		{
			name:    "empty token",
			req:     &v1.GetTaskRequest{},
			token:   "",
			wantErr: true,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return nil, errors.New("verify session error")
				},
			},
		},
		{
			name:    "error from internal-service #0",
			req:     &v1.GetTaskRequest{},
			token:   "token for user-123",
			wantErr: true,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return &togo_user_session_v1.VerifyTokenResponse{
						UserId:   "123",
						Username: "user-123",
					}, nil
				},
			},
			internal: &mock_togo_internal_v1.TogoInternalServiceClientMock{
				GetTaskFunc: func(ctx context.Context, in *togo_internal_v1.GetTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.Task, error) {
					return nil, status.Error(codes.InvalidArgument, codes.Internal.String())
				},
			},
		},
		{
			name:    "error from internal-service #1",
			req:     &v1.GetTaskRequest{},
			token:   "token for user-123",
			wantErr: true,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return &togo_user_session_v1.VerifyTokenResponse{
						UserId:   "123",
						Username: "user-123",
					}, nil
				},
			},
			internal: &mock_togo_internal_v1.TogoInternalServiceClientMock{
				GetTaskFunc: func(ctx context.Context, in *togo_internal_v1.GetTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.Task, error) {
					return nil, errors.New("other error")
				},
			},
		},
		{
			name: "happy case #0",
			req: &v1.GetTaskRequest{
				Id: "task-123",
			},
			token:   "token for user-123",
			wantErr: false,
			want: &v1.Task{
				Id:      "task-123",
				UserId:  "123",
				Title:   "the title",
				Content: "content",
			},
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return &togo_user_session_v1.VerifyTokenResponse{
						UserId:   "123",
						Username: "user-123",
					}, nil
				},
			},
			internal: &mock_togo_internal_v1.TogoInternalServiceClientMock{
				GetTaskFunc: func(ctx context.Context, in *togo_internal_v1.GetTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.Task, error) {
					return &togo_internal_v1.Task{
						Id:      "task-123",
						UserId:  "123",
						Title:   "the title",
						Content: "content",
					}, nil
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{
				"authorization": []string{"Bearer " + test.token},
			})
			authFunc := auth.NewAuthFunc(test.session)
			ctx, err := authFunc(ctx)

			if err != nil {
				assert.True(t, test.wantErr, "GetTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			h := Handler{
				TogoInternalService:    test.internal,
				TogoUserSessionService: test.session,
			}
			resp, err := h.GetTask(ctx, test.req)
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
