package dto

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/domain"
	"net/http"
)

func ResponseError(ctx *gin.Context, err error, statusCode ...int) {
	status := http.StatusInternalServerError
	if errors.Is(err, domain.ErrFailPrecondition) {
		status = http.StatusBadRequest
	}
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	ctx.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = ""
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
