package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"togo-public-api/internal/service/togo_user_session_v1"
	"togo-public-api/mock/mock_togo_user_session_v1"
)

func Test_Auth(t *testing.T) {

	tests := []struct {
		name    string
		mock    togo_user_session_v1.TogoUserSessionServiceClient
		wantErr bool
		want    string
		meta    metadata.MD
	}{
		{
			name:    "empty metadata",
			wantErr: true,
			want:    "",
			meta:    nil,
			mock:    nil,
		},
		{
			name:    "meta data not contained authorization key",
			wantErr: true,
			want:    "",
			meta: metadata.MD{
				"authorization-xxx": []string{},
			},
			mock: nil,
		},
		{
			name:    "empty token #0",
			wantErr: true,
			want:    "",
			meta: metadata.MD{
				"authorization": []string{""},
			},
		},
		{
			name:    "empty token #1",
			wantErr: true,
			want:    "",
			meta: metadata.MD{
				"authorization": []string{"Bearer "},
			},
		},
		{
			name:    "verify token response error",
			wantErr: true,
			want:    "",
			meta: metadata.MD{
				"authorization": []string{"Bearer abc"},
			},
			mock: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return nil, errors.New("internal")
				},
			},
		},
		{
			name:    "happy case #0",
			wantErr: false,
			want:    "1",
			meta: metadata.MD{
				"authorization": []string{"Bearer abc"},
			},
			mock: &mock_togo_user_session_v1.TogoUserSessionServiceClientMock{
				VerifyTokenFunc: func(ctx context.Context, in *togo_user_session_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*togo_user_session_v1.VerifyTokenResponse, error) {
					return &togo_user_session_v1.VerifyTokenResponse{
						UserId:   "1",
						Username: "username-1",
					}, nil
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authFunc := NewAuthFunc(test.mock)
			inCtx := context.TODO()
			inCtx = metadata.NewIncomingContext(inCtx, test.meta)
			newCtx, err := authFunc(inCtx)
			if (err != nil) != test.wantErr {
				t.Errorf("AuthFunc() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			got := GetUserID(newCtx)
			assert.Equal(t, test.want, got)
		})
	}
}
