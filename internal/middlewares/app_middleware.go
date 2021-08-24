package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/helpers"
	"net/http"
	"strings"
)

type AppMiddleware struct {
	tokenProvider helpers.TokenProvider
}

func NewAppMiddleware(injectedTokenProvider helpers.TokenProvider) *AppMiddleware {
	return &AppMiddleware{
		tokenProvider: injectedTokenProvider,
	}
}
func (m *AppMiddleware) ApplyCorsFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		if ctx.Request.Method == http.MethodOptions {
			ctx.Writer.WriteHeader(http.StatusOK)
		}
		ctx.Next()
	}
}

func (m *AppMiddleware) ApplyJwtFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.Request.URL.Path, "login") {
			ctx.Next()
			return
		}

		token := ctx.Request.Header.Get(constants.KeyAuthorization)
		userId, ok := m.tokenProvider.ValidateToken(token)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(constants.KeyUserId, userId)
		ctx.Next()
	}
}
