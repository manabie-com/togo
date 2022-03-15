package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
	UserID  uint
	IsAdmin bool
}

func CreateToken(userid uint, isAdmin bool) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	// Create the Claims
	claims := jwt.MapClaims{
		"userid": userid,
		"admin":  isAdmin,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		userId := uint(claims["userid"].(float64))
		isAdmin := claims["admin"].(bool)

		now := time.Now().Unix()

		// Checking, if now time greather than expiration from JWT.
		if now > expires {
			return nil, fmt.Errorf("unauthorized, token expired")
		}

		return &TokenMetadata{
			Expires: expires,
			UserID:  userId,
			IsAdmin: isAdmin,
		}, nil
	}

	return nil, fmt.Errorf("unauthorized, please login to system")
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
