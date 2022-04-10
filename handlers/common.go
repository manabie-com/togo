package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/namnhatdoan/togo/constants"
	"io"
)

func bindingBodyData(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(&req); err != nil {
		switch {
		case err == io.EOF:
			badRequestError(c, constants.RequestBodyEmpty)
			return false
		default:
			badRequestError(c, err.Error())
			return false
		}
	}
	return true
}

