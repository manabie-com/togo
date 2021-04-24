package http

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func jwtAuthMiddleware(secret []byte) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Request().Header.Get("Authorization")
			claims := make(jwt.MapClaims)
			t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, nil)
			}
			if !t.Valid {
				return ctx.JSON(http.StatusUnauthorized, nil)
			}
			id, ok := claims["user_id"].(string)
			if !ok {
				return ctx.JSON(http.StatusUnauthorized, nil)
			}
			ctx.Set(userAuthKey, id)
			return h(ctx)
		}
	}
}
