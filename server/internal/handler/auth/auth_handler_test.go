package auth

import (
	"context"
	"errors"
	authMock "github.com/HoangVyDuong/togo/internal/usecase/auth"
	"github.com/HoangVyDuong/togo/pkg/define"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"testing"
)

func Test_authHandler_Auth(t *testing.T) {
	viper.Set("jwt.key", "abc")
	trueToken, _ := utils.CreateToken(strconv.FormatInt(1, 10), viper.GetString("jwt.key"))
	tests := []struct {
		name         string
		request authDTO.AuthUserRequest
		authResp int64
		authError error
		wantResponse authDTO.AuthUserResponse
		wantErr      error
	}{
		{
			name: "Success Case",
			request: authDTO.AuthUserRequest{
				Username: "Hoang VY",
				Password: "password",
			},
			authResp: 1,
			wantResponse: authDTO.AuthUserResponse{
				Token: trueToken,
			},
		},
		{
			name: "Unvalidated (username empty)",
			request: authDTO.AuthUserRequest{
				Username: "",
				Password: "password",
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "Unvalidated (password empty)",
			request: authDTO.AuthUserRequest{
				Username: "Hoang Vy",
				Password: "",
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "Auth Failed Not By Notfound or FailPassword",
			authError: errors.New("any error"),
			wantErr: define.Unknown,
		},
		{
			name: "Auth Failed By Notfound",
			authError: define.AccountNotExist,
			wantErr: define.AccountNotExist,
		},
		{
			name: "Auth Failed By FailPassword",
			authError: define.AccountNotAuthorized,
			wantErr: define.AccountNotAuthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authServiceMock := &authMock.ServiceMock{}
			authServiceMock.AuthFunc = func(ctx context.Context, userName string, password string) (int64, error) {
				return tt.authResp, tt.authError
			}
			handler := &authHandler{
				authService: authServiceMock,
			}

			gotResponse, err := handler.Auth(context.Background(), tt.request)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Auth() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}