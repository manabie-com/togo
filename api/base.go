package api

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, success bool, message string, data interface{}, err error) {
	g.C.JSON(httpCode, gin.H{
		"success": success,
		"message": message,
		"data":    data,
		"error":   err,
	})
	return
}

func (g *Gin) BindAndValidate(obj interface{}) bool {
	err := g.C.ShouldBind(obj)
	if err != nil {
		g.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
		return false
	}

	isValid, err := govalidator.ValidateStruct(obj)
	if err != nil || !isValid {
		g.Response(http.StatusBadRequest, false, err.Error(), nil, err)
		return false
	} else {
		return true
	}
}
