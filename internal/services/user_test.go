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

func Test_userService_CreateUser(t *testing.T) {
	type fields struct {
		passwordHashProvider provider.PasswordHashProvider
		userRepo             repository.UserRepository
	}
	type args struct {
		ctx   context.Context
		input *domain.User
	}

	// Mocks
	userId := uint(1)
	fullName := faker.Name()
	username := faker.Username()
	username2 := faker.Username()
	password := faker.Password()

	userInput := &domain.User{
		FullName: fullName,
		Username: username,
		Password: password,
		TasksPerDay: 1,
	}

	userInput2 := &domain.User{
		FullName: fullName,
		Username: username2,
		Password: password,
		TasksPerDay: 1,
	}

	user := &domain.User{
		ID: userId,
		FullName: fullName,
		Username: username,
		Password: password,
		TasksPerDay: 1,
	}

	user2 := &domain.User{
		ID: 2,
		FullName: fullName,
		Username: username2,
		Password: password,
		TasksPerDay: 1,
	}

	passwordHashProvider := new(mockPasswordHashProvider)
	passwordHashProvider.On("HashPassword", mock.Anything).Return(password, nil)
	userRepo := new(mockUserRepository)
	userRepo.On("FindOne", &domain.User{Username: username2}).Return(user2, nil)
	userRepo.On("FindOne", mock.Anything).Return(nil, domain.ErrUserNotFound)
	userRepo.On("Create", userInput).Return(user, nil)
	brokenUserRepo := new(mockUserRepository)
	brokenUserRepo.On("FindOne", mock.Anything).Return(nil, domain.ErrUserNotFound)
	brokenUserRepo.On("Create", mock.Anything).Return(nil, errors.New("invalid"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "username is existed",
			fields: fields{passwordHashProvider, userRepo},
			args: args{
				context.Background(),
				userInput2,
			},
			wantErr: true,
		},
		{
			name: "user create failed",
			fields: fields{passwordHashProvider, brokenUserRepo},
			args: args{
				context.Background(),
				userInput,
			},
			wantErr: true,
		},
		{
			name: "user create successfully",
			fields: fields{passwordHashProvider, userRepo},
			args: args{
				context.Background(),
				userInput,
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				passwordHashProvider: tt.fields.passwordHashProvider,
				userRepo:             tt.fields.userRepo,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByID(t *testing.T) {
	type fields struct {
		passwordHashProvider provider.PasswordHashProvider
		userRepo             repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  uint
	}

	// Mocks
	notFoundUserID := uint(1)
	userId := uint(2)

	user := &domain.User{
		ID: userId,
		FullName: faker.Name(),
		Username: faker.Username(),
		Password: faker.Password(),
		TasksPerDay: 1,
	}
	userRepo := new(mockUserRepository)
	userRepo.On("FindOne", &domain.User{ID: notFoundUserID}).Return(nil, errors.New("USER_NOT_FOUND"))
	userRepo.On("FindOne", &domain.User{ID: userId}).Return(user, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name:   "Not found user",
			fields: fields{userRepo: userRepo},
			args: args{
				context.Background(),
				notFoundUserID,
			},
			wantErr: true,
		},
		{
			name:   "Find user successfully",
			fields: fields{userRepo: userRepo},
			args: args{
				context.Background(),
				userId,
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				passwordHashProvider: tt.fields.passwordHashProvider,
				userRepo:             tt.fields.userRepo,
			}
			got, err := s.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateByID(t *testing.T) {
	type fields struct {
		passwordHashProvider provider.PasswordHashProvider
		userRepo             repository.UserRepository
	}
	type args struct {
		ctx    context.Context
		id     uint
		update *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				passwordHashProvider: tt.fields.passwordHashProvider,
				userRepo:             tt.fields.userRepo,
			}
			got, err := s.UpdateByID(tt.args.ctx, tt.args.id, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.UpdateByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
