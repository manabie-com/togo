package middleware

import (
	"errors"
	"fmt"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/util"

	"github.com/labstack/echo/v4"
)

func checkAuthenticate(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if len(token) <= 0 {
		return errors.New("missing authorization")
	}

	claims := &util.Claims{}
	err := util.DecodeToken(token, claims, config.Cfg.JwtKey)
	if err != nil {
		return err
	}
	c.Set("user", claims)

	return nil
}

// IsAuthenticate func
func IsAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := checkAuthenticate(c); err != nil {
			fmt.Println(err)
			return err
		}
		return next(c)
	}
}

// GetEmailFromContext func
func GetEmailFromContext(c echo.Context) string {
	claims := c.Get("user").(*util.Claims)
	return claims.Email
}

// GetUserIDFromContext func
func GetUserIDFromContext(c echo.Context) uint64 {
	claims := c.Get("user").(*util.Claims)
	return claims.ID
}
