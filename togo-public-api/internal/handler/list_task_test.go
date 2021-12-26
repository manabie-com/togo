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

func Test_ListTask(t *testing.T) {
	tests := []struct {
		name     string
		req      *v1.ListTaskRequest
		token    string
		want     *v1.ListTaskResponse
		wantErr  bool
		internal togo_internal_v1.TogoInternalServiceClient
		session  togo_user_session_v1.TogoUserSessionServiceClient
	}{
		{
			name:    "empty token",
			req:     &v1.ListTaskRequest{},
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
			req:     &v1.ListTaskRequest{},
			token:   "token for user-123",
			want:    nil,
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
				ListTaskFunc: func(ctx context.Context, in *togo_internal_v1.ListTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.ListTaskResponse, error) {
					return nil, status.Error(codes.InvalidArgument, codes.Internal.String())
				},
			},
		},
		{
			name:    "error from internal-service #1",
			req:     &v1.ListTaskRequest{},
			token:   "token for user-123",
			want:    nil,
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
				ListTaskFunc: func(ctx context.Context, in *togo_internal_v1.ListTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.ListTaskResponse, error) {
					return nil, errors.New("other error")
				},
			},
		},
		{
			name:  "happy case #0",
			req:   &v1.ListTaskRequest{},
			token: "token for user-123",
			want: &v1.ListTaskResponse{
				Tasks: []*v1.Task{
					{
						Id:      "id-01",
						UserId:  "user-01",
						Title:   "title-01",
						Content: "content-01",
					},
					{
						Id:      "id-02",
						UserId:  "user-02",
						Title:   "title-02",
						Content: "content-02",
					},
				},
				NextPageToken: "next page",
			},
			wantErr: false,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return &togo_user_session_v1.VerifyTokenResponse{
						UserId:   "123",
						Username: "user-123",
					}, nil
				},
			},
			internal: &mock_togo_internal_v1.TogoInternalServiceClientMock{
				ListTaskFunc: func(ctx context.Context, in *togo_internal_v1.ListTaskRequest, opts ...grpc.CallOption) (*togo_internal_v1.ListTaskResponse, error) {
					return &togo_internal_v1.ListTaskResponse{
						Tasks: []*togo_internal_v1.Task{
							{
								Id:      "id-01",
								UserId:  "user-01",
								Title:   "title-01",
								Content: "content-01",
							},
							{
								Id:      "id-02",
								UserId:  "user-02",
								Title:   "title-02",
								Content: "content-02",
							},
						},
						NextPageToken: "next page",
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
				assert.True(t, test.wantErr, "ListTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			h := Handler{
				TogoInternalService:    test.internal,
				TogoUserSessionService: test.session,
			}
			resp, err := h.ListTask(ctx, test.req)
			if (err != nil) != test.wantErr {
				t.Errorf("ListTask() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !assert.True(t, proto.Equal(resp, test.want)) {
				t.Errorf("ListTask() resp = %v, want %v", resp, test.want)
				return
			}
		})
	}
}
