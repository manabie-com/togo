package middlewares

import (
	"log"

	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/utils"

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
