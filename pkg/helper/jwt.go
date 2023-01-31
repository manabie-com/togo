package helper

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/trinhdaiphuc/togo/internal/entities"
)

// JwtClaims a struct for our custom JWT payload.
type JwtClaims struct {
	UserID   int    `json:"uid"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

var (
	ErrClaimsNotfound = fiber.NewError(fiber.StatusUnauthorized, "User claims not found in context")
	ErrUserNotFound   = fiber.NewError(fiber.StatusUnauthorized, "User not found in context")
)

func SignToken(claims JwtClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func UserFromContext(c context.Context) (*entities.User, error) {
	// Get userID from the previous route.
	jwtData, ok := c.Value("user").(*jwt.Token)
	if !ok {
		return nil, ErrUserNotFound
	}

	claims, ok := jwtData.Claims.(*JwtClaims)
	if !ok {
		return nil, ErrClaimsNotfound
	}

	return &entities.User{
		ID:       claims.UserID,
		Username: claims.UserName,
	}, nil
}
