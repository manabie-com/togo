package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/common/request"
	"github.com/shanenoi/togo/common/response"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/domain"
	"net/http"
)

func JwtValidator(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := request.GetAuthToken(ctx); token != "" {
			if err := domain.NewUserDomain(ctx).CheckUserToken(token); err != nil {
				ctx.JSON(http.StatusUnauthorized, response.Failure(config.RESP_UNAUTHORIZED))
			} else {
				fn(ctx)
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, response.Failure(config.RESP_UNAUTHORIZED))
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
				"Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
