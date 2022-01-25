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

func SignToken(claims JwtClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func UserFromContext(c *fiber.Ctx) (*entities.User, error) {
	// Get userID from the previous route.
	jwtData, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}

	claims, ok := jwtData.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("user claim not found in context")
	}

	return &entities.User{
		ID:       claims.UserID,
		Username: claims.UserName,
	}, nil
}
