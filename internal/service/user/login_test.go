package userservice

import (
	"context"
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/repository/mock"
)

func Test_UserServiceLogin(t *testing.T) {
	type args struct {
		req *entities.User
	}
	var (
		ctx  = context.Background()
		err  = errors.New("got error")
		req  = &entities.User{Username: "daiphuc", Password: "123456"}
		resp = &entities.User{Username: "daiphuc", Password: "$2a$07$oZl47NP//MvMopBDwT3MHuITTNX7T9ENTYwHm1gColLwyymqxayP2"}
	)

	tests := []struct {
		name       string
		args       args
		onExpect   func(userRepo *mock.MockUserRepository)
		expectErr  error
		expectResp *entities.User
	}{
		{
			name: "get user failed",
			args: args{
				req: req,
			},
			onExpect: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().GetUserByName(ctx, req.Username).Return(nil, err)
			},
			expectErr:  err,
			expectResp: nil,
		},
		{
			name: "get user success but password not match",
			args: args{
				req: &entities.User{Username: "daiphuc", Password: "1234567"},
			},
			onExpect: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().GetUserByName(ctx, req.Username).Return(resp, nil)
			},
			expectErr:  fiber.ErrUnauthorized,
			expectResp: nil,
		},
		{
			name: "login user success",
			args: args{
				req: req,
			},
			onExpect: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().GetUserByName(ctx, req.Username).Return(resp, nil)
			},
			expectErr:  nil,
			expectResp: &entities.User{Username: "daiphuc", TaskLimit: 0, ID: 0},
		},
	}

	mockCtrl := gomock.NewController(t)
	userRepo := mock.NewMockUserRepository(mockCtrl)
	cfg := &configs.Config{
		JwtSecret: "secret",
	}
	service := NewUserService(userRepo, cfg)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.onExpect(userRepo)
			response, err := service.Login(ctx, tt.args.req)
			if err != nil || tt.expectErr != nil {
				assert.EqualErrorf(t, tt.expectErr, err.Error(), "Expected: %v, got: %v", tt.expectResp, response)
			}
			if response != nil && tt.expectResp != nil {
				assert.Equal(t, tt.expectResp.Username, response.Username)
				assert.Equal(t, tt.expectResp.TaskLimit, response.TaskLimit)
				assert.Equal(t, tt.expectResp.ID, response.ID)
			}
		})
	}
}
