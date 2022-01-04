package handlers

import (
	"github.com/gin-gonic/gin"
)

type ToGoHandlerI interface {
	CreateTask(c *gin.Context)
	SetConfig(c *gin.Context)
}
