package services

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
)

func Test_userService_GetAuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	id := "firstUser"
	pwd := "example"
	type fields struct {
		repo *repositories.Repository
	}
	type args struct {
		ctx     context.Context
		userReq models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get_auth_token_successfully",
			fields: fields{
				repo: &repositories.Repository{
					UserRepository: func() *repositories.MockUserRepository {
						repo := repositories.NewMockUserRepository(ctrl)
						repo.EXPECT().ValidateUser(ctx, id).Return(
							&models.User{
								ID:       id,
								Password: pwd,
							}, nil,
						)
						return repo
					}(),
				},
			},
			args: args{
				ctx: ctx,
				userReq: models.User{
					ID:       id,
					Password: pwd,
				},
			},
		},
		{
			name: "password_not_match",
			fields: fields{
				repo: &repositories.Repository{
					UserRepository: func() *repositories.MockUserRepository {
						repo := repositories.NewMockUserRepository(ctrl)
						repo.EXPECT().ValidateUser(ctx, id).Return(
							&models.User{
								ID:       id,
								Password: "abcxyz",
							}, nil,
						)
						return repo
					}(),
				},
			},
			args: args{
				ctx: ctx,
				userReq: models.User{
					ID:       id,
					Password: pwd,
				},
			},
			wantErr: true,
		},
		{
			name: "user_not_found",
			fields: fields{
				repo: &repositories.Repository{
					UserRepository: func() *repositories.MockUserRepository {
						repo := repositories.NewMockUserRepository(ctrl)
						repo.EXPECT().ValidateUser(ctx, id).Return(nil, errors.New("user not found"))
						return repo
					}(),
				},
			},
			args: args{
				ctx: ctx,
				userReq: models.User{
					ID:       id,
					Password: pwd,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				repo: tt.fields.repo,
			}
			got, err := s.GetAuthToken(tt.args.ctx, tt.args.userReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" && err == nil {
				t.Errorf("userService.GetAuthToken() = %v", got)
			}
		})
	}
}
