package auth

import (
	"context"
	userEntity "github.com/HoangVyDuong/togo/internal/storages/user"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"reflect"
	"testing"
)

func Test_authService_Auth(t *testing.T) {
	trueHashPassword, _ := utils.GeneratePassword("password")
	tests := []struct {
		name    string
		username string
		password string
		getUserByNameResp userEntity.User
		getUserByNameErr error
		want    uint64
		wantErr error
	}{
		{
			name: "Success Case",
			username: "HoangVyDuong",
			password: "password",
			getUserByNameResp: userEntity.User{
				ID: 111,
				Name: "HoangVyDuong",
				Password: trueHashPassword,
			},
			want: 111,
		},
		{
			name: "user name empty",
			username: "",
			password: "password",
			wantErr: define.FailedValidation,
		},
		{
			name: "password empty",
			username: "HoangVyDuong",
			password: "",
			wantErr: define.FailedValidation,
		},
		{
			name: "Database Error",
			username: "HoangVyDuong",
			password: "password",
			getUserByNameErr: define.DatabaseError,
			wantErr: define.DatabaseError,
		},
		{
			name: "Account Not Fount",
			username: "HoangVyDuong",
			password: "password",
			getUserByNameErr: define.AccountNotExist,
			wantErr: define.AccountNotExist,
		},
		{
			name: "Password Not Match",
			username: "HoangVyDuong",
			password: "password",
			getUserByNameResp: userEntity.User{
				ID: 111,
				Name: "HoangVyDuong",
				Password: "fake-password",
			},
			wantErr: define.AccountNotAuthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := &RepositoryMock{}
			userRepoMock.GetUserByNameFunc = func(ctx context.Context, name string) (userEntity.User, error) {
				return tt.getUserByNameResp, tt.getUserByNameErr
			}
			s := &authService{
				repo: userRepoMock,
			}
			got, err := s.Auth(context.Background(), tt.username, tt.password)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth() got = %v, want %v", got, tt.want)
			}
		})
	}
}