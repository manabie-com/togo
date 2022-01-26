package userservice

import (
	"context"
	"errors"
	"github.com/bmizerany/assert"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/repository"
	"github.com/trinhdaiphuc/togo/internal/repository/mock"
	"testing"
)

func Test_UserServiceLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type fields struct {
		userRepo repository.UserRepository
		cfg      *configs.Config
	}
	type args struct {
		req *entities.User
	}
	var (
		ctx  = context.Background()
		err  = errors.New("got error")
		req  = &entities.User{Username: "daiphuc", Password: "123456"}
		resp = &entities.User{Username: "daiphuc", Password: "123456"}
	)

	tests := []struct {
		name    string
		args    args
		fields  fields
		errResp error
		want    *entities.User
	}{
		{
			name: "get user failed",
			args: args{
				req: req,
			},
			fields: fields{
				userRepo: func() *mock.MockUserRepository {
					mockRepo := mock.NewMockUserRepository(mockCtrl)
					mockRepo.EXPECT().GetUserByName(ctx, req.Username).Return(nil, err)
					return mockRepo
				}(),
			},
			errResp: err,
			want:    nil,
		},
		{
			name: "get user success but password not match",
			args: args{
				req: &entities.User{Username: "daiphuc", Password: "1234567"},
			},
			fields: fields{
				userRepo: func() *mock.MockUserRepository {
					mockRepo := mock.NewMockUserRepository(mockCtrl)
					mockRepo.EXPECT().GetUserByName(ctx, req.Username).Return(resp, nil)
					return mockRepo
				}(),
			},
			errResp: fiber.ErrUnauthorized,
			want:    nil,
		},
		{
			name: "login user success",
			args: args{
				req: req,
			},
			fields: fields{
				userRepo: func() *mock.MockUserRepository {
					mockRepo := mock.NewMockUserRepository(mockCtrl)
					mockRepo.EXPECT().GetUserByName(ctx, req.Username).Return(resp, nil)
					return mockRepo
				}(),
				cfg: &configs.Config{
					JwtSecret: "secret",
				},
			},
			errResp: nil,
			want:    &entities.User{Username: "daiphuc", TaskLimit: 0, ID: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepo: tt.fields.userRepo,
				cfg:      tt.fields.cfg,
			}
			got, errOut := u.Login(ctx, tt.args.req)
			assert.Equal(t, tt.errResp, errOut)
			if got != nil && tt.want != nil {
				assert.Equal(t, tt.want.Username, got.Username)
				assert.Equal(t, tt.want.TaskLimit, got.TaskLimit)
				assert.Equal(t, tt.want.ID, got.ID)
			}
		})
	}
}
