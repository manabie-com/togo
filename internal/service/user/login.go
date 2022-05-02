package userservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"todo/internal/entities"
	"todo/pkg/helper"
	"time"
)

func (u *userService) Login(ctx context.Context, user *entities.User) (*entities.User, error) {
	logrus.Info("Login")
	userResp, err := u.userRepo.GetUserByName(ctx, user.Username)
	if err != nil {
		logrus.Errorf("GetUserByName err %v", err)
		return nil, err
	}
	if helper.CheckPasswordHash(userResp.Password, user.Password) {
		return nil, fiber.ErrUnauthorized
	}

	// Create token
	jwtClaim := helper.JwtClaims{
		UserID:   userResp.ID,
		UserName: userResp.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	userResp.Jwt, err = helper.SignToken(jwtClaim, u.cfg.JwtSecret)
	if err != nil {
		logrus.Errorf("SignToken err %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	userResp.Password = ""
	return userResp, nil
}
