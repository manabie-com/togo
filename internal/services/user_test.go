package services

import (
	"context"
	"reflect"
	"testing"
	"togo/internal/domain"
	"togo/internal/provider"
	"togo/internal/repository"
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
