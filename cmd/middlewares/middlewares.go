package middlewares

import (
	"log"

	"example.com/m/v2/constants"
	"example.com/m/v2/internal/api/handlers"
	"example.com/m/v2/utils"

	"github.com/gin-gonic/gin"
)

func SetDefaultMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")

		ctx.Header("Content-Type", "application/json")
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, err := handlers.GetValueCookieFromCtx(ctx, constants.CookieTokenKey)
		if err != nil {
			log.Fatalf("Fail GetTokenKeyFromCtx")
		}

		if utils.SafeString(value) == "" {
			log.Fatalf("Please login")
		}
	}
}
