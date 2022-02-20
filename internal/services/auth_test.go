package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"togo/internal/domain"
	"togo/internal/provider"
	"togo/internal/repository"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
)

func Test_authService_Login(t *testing.T) {
	type fields struct {
		passwordHashProvider provider.PasswordHashProvider
		tokenProvider        provider.TokenProvider
		userRepo             repository.UserRepository
	}
	type args struct {
		ctx        context.Context
		credential *domain.LoginCredential
	}
	// Mocks
	userID := uint(1)
	fullName := faker.Name()
	username := faker.Username()
	password := faker.Password()
	passwordHash := faker.Password()
	user := &domain.User{
		ID:          userID,
		FullName:    fullName,
		Username:    username,
		Password:    passwordHash,
		TasksPerDay: 1,
	}
	jwtToken := faker.Jwt()
	userRepo := new(mockUserRepository)
	userRepo.On("FindOne", mock.Anything).Return(user, nil)
	userRepoUserNotFound := new(mockUserRepository)
	userRepoUserNotFound.On("FindOne", mock.Anything).Return(nil, domain.ErrUserNotFound)
	passwordHashProvider := new(mockPasswordHashProvider)
	passwordHashProvider.On("HashPassword", mock.Anything).Return(passwordHash, nil)
	passwordHashProvider.On("ComparePassword", password, passwordHash).Return(nil)
	passwordHashProviderComparePasswordFailed := new(mockPasswordHashProvider)
	passwordHashProviderComparePasswordFailed.On("ComparePassword", password, passwordHash).Return(errors.New("error"))
	tokenProvider := new(mockTokenProvider)
	tokenProvider.On("GenerateToken", user).Return(jwtToken, nil)
	tokenProviderGenerateTokenFailed := new(mockTokenProvider)
	tokenProviderGenerateTokenFailed.On("GenerateToken", user).Return("", errors.New("error"))
	// Test cases
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.LoginResult
		wantErr bool
	}{
		{
			name: "user not exists",
			fields: fields{
				userRepo: userRepoUserNotFound,
			},
			args: args{
				context.Background(),
				&domain.LoginCredential{Username: username, Password: password},
			},
			wantErr: true,
		},
		{
			name: "password incorrect",
			fields: fields{
				userRepo:             userRepo,
				passwordHashProvider: passwordHashProviderComparePasswordFailed,
			},
			args: args{
				context.Background(),
				&domain.LoginCredential{Username: username, Password: password},
			},
			wantErr: true,
		},
		{
			name: "generate token failed",
			fields: fields{
				userRepo:             userRepo,
				passwordHashProvider: passwordHashProvider,
				tokenProvider:        tokenProviderGenerateTokenFailed,
			},
			args: args{
				context.Background(),
				&domain.LoginCredential{Username: username, Password: password},
			},
			wantErr: true,
		},
		{
			name: "login successfully",
			fields: fields{
				userRepo:             userRepo,
				passwordHashProvider: passwordHashProvider,
				tokenProvider:        tokenProvider,
			},
			args: args{
				context.Background(),
				&domain.LoginCredential{Username: username, Password: password},
			},
			want: &domain.LoginResult{
				Profile: user,
				Token:   jwtToken,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := authService{
				passwordHashProvider: tt.fields.passwordHashProvider,
				tokenProvider:        tt.fields.tokenProvider,
				userRepo:             tt.fields.userRepo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.credential)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_VerifyToken(t *testing.T) {
	type fields struct {
		passwordHashProvider provider.PasswordHashProvider
		tokenProvider        provider.TokenProvider
		userRepo             repository.UserRepository
	}
	type args struct {
		ctx   context.Context
		token string
	}
	// Mocks
	userID := uint(1)
	payload := &domain.User{ID: userID}
	user := &domain.User{
		ID:          userID,
		FullName:    faker.Name(),
		Username:    faker.Username(),
		Password:    faker.Password(),
		TasksPerDay: 1,
	}
	token := faker.Jwt()
	userRepo := new(mockUserRepository)
	userRepo.On("FindOne", mock.Anything).Return(user, nil)
	userRepoUserNotFound := new(mockUserRepository)
	userRepoUserNotFound.On("FindOne", mock.Anything).Return(nil, domain.ErrUserNotFound)
	tokenProvider := new(mockTokenProvider)
	tokenProvider.On("VerifyToken", mock.Anything).Return(payload, nil)
	tokenProviderVerifyTokenFailed := new(mockTokenProvider)
	tokenProviderVerifyTokenFailed.On("VerifyToken", mock.Anything).Return(nil, errors.New("error"))
	// Test cases
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.VerifyTokenResult
		wantErr bool
	}{
		{
			name: "token invalid",
			fields: fields{
				tokenProvider: tokenProviderVerifyTokenFailed,
			},
			args: args{
				context.Background(),
				token,
			},
			wantErr: true,
		},
		{
			name: "user not found",
			fields: fields{
				tokenProvider: tokenProvider,
				userRepo:      userRepoUserNotFound,
			},
			args: args{
				context.Background(),
				token,
			},
			wantErr: true,
		},
		{
			name: "verify token successfully",
			fields: fields{
				tokenProvider: tokenProvider,
				userRepo:      userRepo,
			},
			args: args{
				context.Background(),
				token,
			},
			want: &domain.VerifyTokenResult{
				Authenticated: true,
				Payload:       user,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := authService{
				passwordHashProvider: tt.fields.passwordHashProvider,
				tokenProvider:        tt.fields.tokenProvider,
				userRepo:             tt.fields.userRepo,
			}
			got, err := s.VerifyToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
