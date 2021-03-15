package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/jwtauth/v5"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/manabie-com/togo/internal/todo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_ValidateUser(t *testing.T) {
	tokenAuth := jwtauth.New("HS256", []byte("fortestonly"), nil)
	jwtExpiredPeriod := 10
	os.Setenv("JWT_EXPIRED_PERIOD", strconv.Itoa(jwtExpiredPeriod))
	assert := assert.New(t)

	tests := []struct {
		name     string
		userRepo d.UserRepository
		cred     d.UserAuthParam
		want     string
		wantErr  bool
	}{
		{
			"System Error",
			func() d.UserRepository {
				repo := &mocks.UserRepository{}
				repo.On("GetByCredentials", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("System Error"))
				return repo
			}(),
			d.UserAuthParam{Username: "test", Password: "test"},
			"",
			true,
		},
		{
			"No valid user",
			func() d.UserRepository {
				repo := &mocks.UserRepository{}
				repo.On("GetByCredentials", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
				return repo
			}(),
			d.UserAuthParam{Username: "test1", Password: "test1"},
			"",
			false,
		},
		{
			"Valid user",
			func() d.UserRepository {
				repo := &mocks.UserRepository{}
				repo.On("GetByCredentials", mock.Anything, mock.Anything, mock.Anything).Return(&d.User{ID: 10}, nil)
				return repo
			}(),
			d.UserAuthParam{Username: "test2", Password: "test2"},
			func() string {
				claims := map[string]interface{}{
					"userID": 10,
					"exp":    time.Now().Add(time.Minute * time.Duration(jwtExpiredPeriod)).Unix(),
				}

				_, token, _ := tokenAuth.Encode(claims)
				return token
			}(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				UserRepo: tt.userRepo,
			}

			ctx := context.Background()
			got, err := s.ValidateUser(ctx, tokenAuth, tt.cred)
			mockRepo, _ := s.UserRepo.(*mocks.UserRepository)

			mockRepo.AssertExpectations(t)
			assert.Equal(tt.want, got)
			assert.True((err != nil) == tt.wantErr)
		})
	}
}
