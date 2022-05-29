package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"togo/domain/model"
	"togo/domain/service"
)

func AuthMiddleware(tokenService service.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.GetHeader("bearer")
		if bearer == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}
		userClaim, err := tokenService.ValidateToken(context.Background(), bearer)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		currentTime := time.Now()
		if userClaim.ExpiresAt < currentTime.UnixMilli() {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"rc": 10,
				"rd": "Relogin",
			})
		}
		ctx.Set("userInfo", model.User{
			Id:    userClaim.UserId,
			Limit: userClaim.Limit,
		})
		ctx.Next()
	}
}
