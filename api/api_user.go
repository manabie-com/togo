package api

import (
	"net/http"
	"togo/model"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	username := c.Query("user_id")
	password := c.Query("password")

	token, err := model.Login(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
