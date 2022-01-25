package helper

import (
	"fmt"
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
	ErrUserNotFound      = fmt.Errorf("user not found in context")
	ErrUserClaimNotFound = fmt.Errorf("user claims not found in context")
)

func SignToken(claims JwtClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func UserFromContext(c *fiber.Ctx) (*entities.User, error) {
	// Get userID from the previous route.
	jwtData, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, ErrUserNotFound
	}

	claims, ok := jwtData.Claims.(*JwtClaims)
	if !ok {
		return nil, ErrUserClaimNotFound
	}

	return &entities.User{
		ID:       claims.UserID,
		Username: claims.UserName,
	}, nil
}
