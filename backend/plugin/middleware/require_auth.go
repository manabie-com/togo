package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/user/usermodel"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func ErrWrongAuthHeader(err error) *sdkcm.AppError {
	return sdkcm.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func RequireAuth(authStore AuthenStore, serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		tokenProvider := serviceCtx.MustGet("jwt").(tokenprovider.Provider)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			panic(err)
		}

		simpleUser := sdkcm.SimpleUser{
			SQLModel: *sdkcm.NewUpsertSQLModel(user.ID),
			Email:    user.Email,
		}
		c.Set("current_user", &simpleUser)
		c.Next()
	}
}
