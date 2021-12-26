package handler

import (
	"context"
	"errors"
	"testing"

	"togo-public-api/internal/service/togo_user_session_v1"
	"togo-public-api/mock/mock_togo_user_session_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func Test_RegisterOrLogin(t *testing.T) {
	tests := []struct {
		name    string
		req     *v1.RegisterOrLoginRequest
		want    *v1.RegisterOrLoginResponse
		wantErr bool
		session togo_user_session_v1.TogoUserSessionServiceClient
	}{
		{
			name:    "error from user-session-service#0",
			req:     &v1.RegisterOrLoginRequest{},
			wantErr: true,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				RegisterOrLoginFunc: func(ctx context.Context, in *togo_user_session_v1.RegisterOrLoginRequest, opts ...grpc.CallOption) (*togo_user_session_v1.RegisterOrLoginResponse, error) {
					return nil, status.Error(codes.InvalidArgument, "")
				},
			},
		},
		{
			name:    "error from user-session-service#1",
			req:     &v1.RegisterOrLoginRequest{},
			wantErr: true,
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				RegisterOrLoginFunc: func(ctx context.Context, in *togo_user_session_v1.RegisterOrLoginRequest, opts ...grpc.CallOption) (*togo_user_session_v1.RegisterOrLoginResponse, error) {
					return nil, errors.New("anything")
				},
			},
		},
		{
			name: "happy case #0",
			req: &v1.RegisterOrLoginRequest{
				Username: "abc",
				Password: "123",
			},
			wantErr: false,
			want: &v1.RegisterOrLoginResponse{
				Token: "token-abc-123",
			},
			session: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				RegisterOrLoginFunc: func(ctx context.Context, in *togo_user_session_v1.RegisterOrLoginRequest, opts ...grpc.CallOption) (*togo_user_session_v1.RegisterOrLoginResponse, error) {
					return &togo_user_session_v1.RegisterOrLoginResponse{
						Token: "token-abc-123",
					}, nil
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			h := Handler{
				TogoUserSessionService: test.session,
			}
			resp, err := h.RegisterOrLogin(ctx, test.req)
			if (err != nil) != test.wantErr {
				t.Errorf("RegisterOrLogin() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !assert.True(t, proto.Equal(resp, test.want)) {
				t.Errorf("RegisterOrLogin() resp = %v, want %v", resp, test.want)
				return
			}
		})
	}
}
