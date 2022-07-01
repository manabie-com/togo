package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/pkg/responses"
	"github.com/manabie-com/togo/utils"
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
			responses.ResponseForError(ctx, err, http.StatusUnauthorized, "Fail Parse token")
			return
		}

		if utils.SafeString(value) == "" {
			responses.ResponseForError(ctx, err, http.StatusUnauthorized, "Not have token")
			return
		}
	}
}
