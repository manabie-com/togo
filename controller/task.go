package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
)

func GetAllTask(c *gin.Context) {
	var task []model.Task
	db.DB.Find(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func GetTaskByID(c *gin.Context) {
	var task model.Task
	if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateTask(c *gin.Context) {
	var input form.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := model.Task{
		Model:   gorm.Model{},
		Content: input.Content,
		UserID:  input.UserID,
	}
	db.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func UpdateTaskByID(c *gin.Context) {
	var task model.Task
	if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTaskByID(c *gin.Context) {
	var task model.Task
	if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
