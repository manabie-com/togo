package gingo

import (
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/gingo/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, ctxErr := range ctx.Errors {
			err := ctxErr.Err

			if ve, ok1 := err.(validator.ValidationErrors); ok1 {
				out := make([]validation.ErrorMsg, len(ve))
				for i, fe := range ve {

					out[i] = validation.ErrorMsg{
						Field:   fe.Field(),
						Message: validation.GetErrorMsg(fe),
					}
				}

				ctx.AbortWithStatusJSON(http.StatusBadRequest, errs.ErrorResponse{
					Code:    errs.ErrCodeValidationError,
					Message: errs.ErrValidationError,
					Errors:  out,
				})
			} else if re, ok2 := err.(*errs.ReadableError); ok2 {
				ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errs.ErrorResponse{
					Code:    re.ErrorCode(),
					Message: re.Message(),
					Error:   re.Message(),
				})
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrorResponse{
					Code:    errs.ErrCodeInternalAppError,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				})
			}
		}
	}
}
