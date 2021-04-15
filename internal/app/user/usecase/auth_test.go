package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/app/user/usecase/mock"
	"testing"
	"time"
)

func TestAuthService_GetAuthToken(t *testing.T) {
	tests := []struct {
		name                string
		mockUserStorage     func() UserStorage
		mockGenerateTokenFn func()
		want                string
		wantErr             error
	}{
		{
			name: "invalid user ID or pwd",
			mockUserStorage: func() UserStorage {
				us := mock.NewMockUserStorage(gomock.NewController(t))
				us.EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error")).AnyTimes()
				return us
			},
			mockGenerateTokenFn: func() {},
			want:                "",
			wantErr:             ErrInvalidUser,
		},
		{
			name: "unable to generate token",
			mockUserStorage: func() UserStorage {
				us := mock.NewMockUserStorage(gomock.NewController(t))
				us.EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return us
			},
			mockGenerateTokenFn: func() {
				generateToken = func(jwtKey, userID string, expiredDuration time.Duration) (string, error) {
					return "", errors.New("mock error")
				}
			},
			want:    "",
			wantErr: ErrUnableToGenerateToken,
		},
		{
			name: "success",
			mockUserStorage: func() UserStorage {
				us := mock.NewMockUserStorage(gomock.NewController(t))
				us.EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return us
			},
			mockGenerateTokenFn: func() {
				generateToken = func(jwtKey, userID string, expiredDuration time.Duration) (string, error) {
					return "mock_token", nil
				}
			},
			want:    "mock_token",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockGenerateTokenFn()
			a := AuthService{
				userStorage: tt.mockUserStorage(),
			}
			got, err := a.GetAuthToken(context.Background(), "", "")
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAuthToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
