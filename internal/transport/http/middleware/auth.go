package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/transport/http/response"
	"github.com/manabie-com/togo/pkg/token"
	"net/http"
	"strings"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayload    = "authorization_payload"
)

type AuthMiddleWare struct {
	tokenMaker token.Token
}

func NewAuthMiddleware(tokenMaker token.Token) AuthMiddleWare {
	return AuthMiddleWare{
		tokenMaker: tokenMaker,
	}
}

func (a AuthMiddleWare) UseAuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := ctx.Request().Header.Get(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "Unauthorized",
			})
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "Unauthorized",
			})
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "unsupported authorization type " + authorizationType,
			})
		}
		accessToken := fields[1]
		payload, err := a.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: err.Error(),
			})
		}
		ctx.Set(AuthorizationPayload, payload)
		return next(ctx)
	}
}
