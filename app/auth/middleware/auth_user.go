package middleware

import (
	"errors"
	"github.com/ansidev/togo/auth/constant"
	"github.com/ansidev/togo/auth/dto"
	authService "github.com/ansidev/togo/auth/service"
	"github.com/ansidev/togo/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthUser(authService authService.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cred := dto.TokenCredential{}

		if err := ctx.ShouldBindHeader(&cred); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse{
				Code:    errs.ErrCodeTokenIsRequired,
				Message: errs.ErrTokenIsRequired,
				Error:   "Authorize token is required",
			})
			return
		}

		token := strings.Split(cred.Token, "Bearer ")

		if len(token) < 2 {
			err := errors.New(errs.ErrInvalidAuthorizationHeader)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errs.ErrorResponse{
				Code:    errs.ErrCodeInvalidAuthorizationHeader,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			return
		}

		authenticationCredential, err := authService.GetCredential(token[1])

		if err != nil {
			errorCode := errs.ErrorCode(err)
			if errorCode == errs.ErrCodeTokenNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse{
					Code:    errs.ErrCodeUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
					Error:   err.Error(),
				})
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrorResponse{
					Code:    errorCode,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				})
			}
			return
		}

		ctx.Set(constant.AuthKey, authenticationCredential)

		ctx.Next()
	}
}
