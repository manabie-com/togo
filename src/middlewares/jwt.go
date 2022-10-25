package middlewares

import (
	"os"
	"strings"
	"todo-api/src/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Jwt() echo.MiddlewareFunc {
	secret := os.Getenv("API_SECRET")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret),
		SigningMethod: "HS256",
		ContextKey:    "token",
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
		Claims:        &models.JwtClaims{},
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Path(), "/login") {
				return true
			}
			return false
		},
	})
}
