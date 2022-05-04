package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"todo/pkg/helper"
)

// JWT error message.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "error",
			"error":   "Missing or malformed JWT!",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"message": "error",
		"error":   "Invalid or expired JWT!",
	})
}

// JWTMiddleware Guards a specific endpoint in the API.
func JWTMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    []byte(secret),
		Claims: 	   &helper.JwtClaims{},
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}
