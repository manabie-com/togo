package tasks

import (
	"net/http"
	"togo/src/modules/users"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Create(c *gin.Context) {

	// validate should split other file
	type RequestBody struct {
		Content string `json:"content" binding:"required" validate:"required"`
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
	user := c.MustGet("user").(users.User)
	task := Task{Content: requestBody.Content, CreatedBy: user}
	result := db.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, task.ToJSON())
}
