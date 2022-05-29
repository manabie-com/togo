package response

import (
	"errors"
	"net/http"

	cerr "github.com/dinhquockhanh/togo/internal/pkg/errors"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(c *gin.Context, err error) {
	res := &cerr.Error{
		Code:    http.StatusInternalServerError,
		Message: "An internal error has occurred. Please contact technical support.",
	}

	defer func() {
		c.JSON(res.Code, res)
	}()

	if err != nil {
		var ae *cerr.Error
		if errors.As(err, &ae) {
			res = ae
			return
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			res = cerr.MakeValidationError(ve)
			return
		}
	}

	// log internal error
	log.WithCtx(c).Errorf("internal err: %v", err)

}
