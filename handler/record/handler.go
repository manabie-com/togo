package record

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	RecordTask(userId, task string) error
}

type Handler struct {
	recordService Service
}

func NewHandler(recordService Service) *Handler {
	return &Handler{
		recordService: recordService,
	}
}

func (h *Handler) HandleRecordTask(c *gin.Context) {
	var reqBody Request
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err := h.recordService.RecordTask(reqBody.UserId, reqBody.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("record user's record error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "record success",
	})
	return
}
