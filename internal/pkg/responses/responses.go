package responses

import (
	"example.com/m/v2/constants"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Status      int
	Error       error
	ErrorDetail string
}

type ResponseOK struct {
	Status int
	Data   interface{}
}

func ResponseForError(ctx *gin.Context, err error, status int, errorDetail string) {
	e := ""
	if err != nil {
		e = err.Error()
	}
	ctx.JSON(status, gin.H{
		constants.ResponseStatus:      status,
		constants.ResponseError:       e,
		constants.ResponseErrorDetail: errorDetail,
	})
}

func ResponseForOK(ctx *gin.Context, status int, data interface{}, message string) {
	ctx.JSON(status, gin.H{
		constants.ResponseStatus:  status,
		constants.ResponseData:    data,
		constants.ResponseMessage: message,
	})
}
