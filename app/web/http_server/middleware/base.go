package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/app/utils/token"

	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// CallFromPublicRoute set "from_public_route":true, detect user call api via public route
func CallFromPublicRoute() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("from_public_route", true) // bool
			return next(c)
		}
	}
}

// ValidateToken check JWT token from header.token
func ValidateToken(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)
			if authorizationHeader == "" {
				return gHandler.NewHTTPError(http.StatusUnauthorized, "token is required", gErrcode.TokenIsRequired)
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				return gHandler.NewHTTPError(http.StatusUnauthorized, "token is invalid", gErrcode.TokenIsInvalid)
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				return gHandler.NewHTTPError(http.StatusUnauthorized, "token is invalid", gErrcode.TokenIsInvalid)
			}

			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				return gHandler.NewHTTPError(http.StatusUnauthorized, "token is invalid", gErrcode.TokenIsExpired)
			}

			c.Set(authorizationPayloadKey, payload)
			return next(c)
		}
	}
}
