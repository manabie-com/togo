package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordService interface {
	RecordTask(userId, task string) error
}

type Handler struct {
	recordService RecordService
}

func NewHandler(recordService RecordService) *Handler {
	return &Handler{
		recordService: recordService,
	}
}

func (h *Handler) HandleRecordTask(c *gin.Context) {
	var reqBody RecordRequest
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err := h.recordService.RecordTask(reqBody.UserId, reqBody.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("record user's task error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "record success",
	})
	return
}
