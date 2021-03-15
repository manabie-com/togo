package auth

import (
	"net/http"

	Users "togo/src/modules/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"togo/src/common/bcrypt"
	"togo/src/common/jwt"
	"togo/src/common/types"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Username string `json:"username" binding:"required" validate:"required"`
		Password string `json:"password" binding:"required" validate:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user Users.User

	// Get first matched record
	if errorQuery := db.Where("username = ?", requestBody.Username).First(&user).Error; errorQuery != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password invalid"})
		return
	}

	hashPassword := user.Password

	match := bcrypt.CheckPasswordHash(requestBody.Password, hashPassword)

	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password invalid"})
		return
	}

	userJson := user.ToJSON()
	token, _ := jwt.GenerateTokenUser(userJson)

	c.JSON(200, types.JSON{
		"user":  userJson,
		"token": token,
	})
}

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
}
