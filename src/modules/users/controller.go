package users

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Create(c *gin.Context) {
	// validate should split other file
	type RequestBody struct {
		Username string `json:"username" binding:"required" validate:"required"`
		Password string `json:"password" binding:"required" validate:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate = validator.New()
	// if errs, ok := validate.Struct(requestBody).(validator.ValidationErrors); ok {
	// 	if errs != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
	// 		return
	// 	}
	// }
	db := c.MustGet("db").(*gorm.DB)
	user := User{Username: requestBody.Username, Password: requestBody.Password}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, user.ToJSON())
}
