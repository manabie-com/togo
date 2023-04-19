// Package usergrp maintains the group of handlers for user access.
package usergrp

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/features/authenticate"
	"github.com/manabie-com/togo/internal/web/auth"
	"github.com/manabie-com/togo/platform/validate"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Auth                *auth.Auth
	AuthenticateFeature *authenticate.Feature
}

// Authenticate provides an API token for the authenticated user.
func (h *Handlers) Authenticate(c echo.Context) error {
	var payload AppAuthenticatingUser
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).WithInternal(err)
	}

	if err := validate.Check(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).WithInternal(err)
	}

	addr, err := mail.ParseAddress(payload.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).WithInternal(err)
	}

	ctx := c.Request().Context()

	usr, err := h.AuthenticateFeature.Authenticate(ctx, *addr, payload.Password)
	if err != nil {
		switch {
		case errors.Is(err, authenticate.ErrNotFound):
			return echo.NewHTTPError(http.StatusNotFound).WithInternal(err)
		case errors.Is(err, authenticate.ErrAuthenticationFailure):
			return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError).WithInternal(err)
		}
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   usr.ID.String(),
			Issuer:    "todo api project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	var tkn struct {
		Token string `json:"token"`
	}

	tkn.Token, err = h.Auth.GenerateToken(claims)
	if err != nil {
		return fmt.Errorf("generatetoken: %w", err)
	}

	return c.JSON(http.StatusOK, tkn)
}
