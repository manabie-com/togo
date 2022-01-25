package userservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/pkg/helper"
	"time"
)

func (u *userService) Login(ctx *fiber.Ctx, user *entities.User) (*entities.User, error) {
	userResp, err := u.userRepo.GetUserByName(ctx.Context(), user.Username)
	if err != nil {
		return nil, err
	}
	if userResp.Username != user.Username || userResp.Password != user.Password {
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
		return nil, err
	}
	userResp.Password = ""
	return userResp, nil
}
