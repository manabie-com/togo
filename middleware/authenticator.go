package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	appctx "github.com/manabie-com/togo/app_ctx"
	"github.com/manabie-com/togo/shared"
	"github.com/manabie-com/togo/user/storage"
	"net/http"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", shared.ErrInvalidRequest(errors.New("invalid token header"))
	}

	return parts[1], nil
}

func Authenticator(appCtx appctx.IAppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenProvider := appCtx.GetTokenProvider()
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		db := appCtx.GetDbConn()
		store := storage.NewUserStorage(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		user, err := store.FindById(c.Request.Context(), payload.GetUserId())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		c.Set("current_user", user)
		c.Next()
	}
}
