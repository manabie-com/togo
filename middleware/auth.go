package middleware

import (
	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/config"
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

func Auth(repo *repository.Repository) func(next echo.HandlerFunc) echo.HandlerFunc {
	handlerFunc := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				ctx   = &util.CustomEchoContext{Context: c}
				token = c.Request().Header.Get("Authorization")
			)

			claims, err := jwt.DecodeToken(token, config.GetConfig().TokenSecretKey)
			if err != nil {
				return util.Response.Error(c, myerror.ErrUnauthorized())
			}

			myUser, err := repo.User.GetByID(ctx, claims.UserID)
			if err != nil {
				return util.Response.Error(c, myerror.ErrUnauthorized())
			}

			c.Set(jwt.MyUserClaim, myUser)

			return next(c)
		}
	}

	return handlerFunc
}
