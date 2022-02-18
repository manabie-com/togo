package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kier1021/togo/api/apierrors"
)

func validationBLMessage(bl apierrors.BLError) []ApiError {
	return []ApiError{
		{
			Param:        bl.Param(),
			ErrorMessage: bl.Error(),
		},
	}
}

// validationMessageForTag sets the validation message depending on the validate struct tag value
func validationMessageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "not_empty":
		return "This field is required"
	case "max":
		return "Value should not be greater than " + fe.Param()
	case "yyyymmdd_date":
		return "Value should be in the YYYY-MM-DD format"
	default:
		return "validation error"
	}
}

// makeErrResponse set the error response to context
func makeErrResponse(err error, c *gin.Context) {

	// Check if the err as validator.ValidationErrors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs := make([]ApiError, len(ve))
		for i, fe := range ve {
			errs[i] = ApiError{fe.Field(), validationMessageForTag(fe)}
		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "Validation Error.",
			"errors":  errs,
		})
		return
	}

	// Check if the err is apierrors.BLError
	var bl apierrors.BLError
	if errors.As(err, &bl) {
		errs := validationBLMessage(bl)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "Validation Error.",
			"errors":  errs,
		})
		return
	}

	// Set the error as internal server error
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal server error occurred.",
		"error":   err.Error(),
	})
}
