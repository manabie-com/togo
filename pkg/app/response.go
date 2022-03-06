package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// Response setting gin.JSON
func (g *Gin) Response(c *gin.Context, httpCode int, status string, errCode int, data interface{}) {
	var message string
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.JSON(httpCode, Response{
		Code:    httpCode,
		Data:    data,
		Error:   message,
		Message: message,
		Status:  status,
	})
	return
}

func (g *Gin) ResponseError(httpCode int, message string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Status:  "error",
		Error:   message,
		Message: message,
		Data:    data,
	})
	return
}

func (g *Gin) SimpleResponse(c *gin.Context, httpCode int, data interface{}) {
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.JSON(httpCode, struct {
		Data interface{} `json:"data"`
	}{Data: data})
	return
}
