package middlewares

import (
	"context"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/manabie-com/togo/modules/auth"
)

var authService auth.Service

// NewAuthenticator return authenticator middleware
func NewAuthenticator(service auth.Service) interface{} {
	authService = service
	return Authenticator
}

// Authenticator is a middleware that authenticate user via JWT token
func Authenticator(ctx *fiber.Ctx) error {
	if checkIfByPassAuthenticating(ctx.Path()) {
		return ctx.Next()
	}

	bearerToken := ctx.Get("Authorization")
	if len := len(bearerToken); len < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Token!")
	}
	userAuth, valid := validToken(bearerToken)
	if valid {
		ctx.Locals("userAuth", userAuth)
		return ctx.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid token!")
}

func checkIfByPassAuthenticating(path string) bool {
	nonAuthRouteStr := "/users/login,/login"

	routes := strings.Split(nonAuthRouteStr, ",")

	for _, route := range routes {
		trimString := strings.TrimSpace(route)
		if trimString == "" {
			break
		}

		result, err := regexp.MatchString(trimString, path)

		if err != nil {
			break
		}

		if result {
			return true
		}
	}

	return false
}

func CreateToken(id string, maxTodo int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["max_todo"] = maxTodo
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(authService.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func validToken(token string) (auth.UserAuth, bool) {
	var userAuth auth.UserAuth
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(authService.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return userAuth, false
	}

	if !t.Valid {
		return userAuth, false
	}

	max_todo, ok2 := claims["max_todo"].(float64)
	id, ok := claims["user_id"].(string)

	//fmt.Println(max_todo)
	if !ok || !ok2 {
		return userAuth, false
	}

	userAuth.UserID = id
	userAuth.MaxTodo = int(max_todo)
	return userAuth, true
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
