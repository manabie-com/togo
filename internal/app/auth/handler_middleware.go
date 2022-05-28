package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
	"github.com/dinhquockhanh/togo/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	authContextKey          = "auth_user"
	authHeaderKey           = "Authorization"
	authorizationTypeBearer = "bearer"
)

func SetUserMiddleware(tokenizer token.Tokenizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader(authHeaderKey)
		if key == "" {
			ctx.Next()
			return
		}
		fields := strings.Fields(key)
		if len(fields) < 2 {
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			return
		}

		payload, err := tokenizer.Verify(fields[1])
		if err != nil {
			log.WithCtx(ctx.Request.Context()).Errorf("verify token: %v", err)
			ctx.Next()
			return
		}
		newCtx := NewContext(ctx.Request.Context(), user.PlayLoadToUser(payload))
		ctx.Request = ctx.Request.WithContext(newCtx)

		ctx.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if usr := FromContext(ctx.Request.Context()); usr == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.Error{
				Code:    http.StatusUnauthorized,
				Message: "missing or invalid authorization header",
			})
			return
		}
	}
}

func NewContext(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, authContextKey, user)
}

// FromContext extract user information from the given context
func FromContext(ctx context.Context) *user.User {
	if v, ok := ctx.Value(authContextKey).(*user.User); ok {
		return v
	}
	return nil
}
