package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/namnhatdoan/togo/constants"
	"net/http"
)

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: constants.Success,
		Msg: constants.SuccessMessage,
		Data: data,
	})
}

func badRequestError(c *gin.Context, err string) {
	response := Response{
		Code: constants.BadRequest,
		Msg: constants.BadRequestMessage,
		Data: err,
	}
	c.JSON(http.StatusOK, &response)
}

func serverError(c *gin.Context, err string) {
	c.JSON(http.StatusOK, &Response{
		Code: constants.GeneralFailure,
		Msg: constants.GeneralFailureMessage,
		Data: err,
	})
}
