package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"togo/config"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func NewRouter(sc *config.ServiceContext) error {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", ping)

	addr := fmt.Sprintf(":%d", sc.Port)
	if err := r.Run(addr); err != nil {
		return err
	}

	return nil
}
