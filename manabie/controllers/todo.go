package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"manabie/manabie/databases/drivers/sqlite"
	libs "manabie/manabie/helpers"
	"manabie/manabie/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTask(c *gin.Context) {
	defer libs.RecoverError(c)
	var (
		status       = 200
		msg          string
		responseData = gin.H{}
		taskModels   []models.Task
	)
	db := sqlite.Connect()
	// filter
	token := c.Request.Header.Get("Authorization")
	userID := libs.GetUserIDFromToken(token)
	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	vCreatedDate, sCreatedDate := libs.GetQueryParam("created_date", c)
	if sCreatedDate {
		db = db.Where("created_date = ?", vCreatedDate)
	}
	result := db.Debug().Find(&taskModels)
	if result.Error != nil {
		status = 500
		msg = result.Error.Error()
	}
	if status == 200 {
		msg = "Success"
		responseData = gin.H{
			"status": status,
			"data":   taskModels,
			"msg":    msg,
		}
	} else {
		if msg == "" {
			msg = "Error"
		}
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}
	libs.APIResponseData(c, status, responseData)
}

func CreateTask(c *gin.Context) {
	defer libs.RecoverError(c)
	var (
		status       = 200
		msg          string
		responseData = gin.H{}
		taskModel    models.Task
		userModel    models.User
		taskModels   []models.Task
	)
	db := sqlite.Connect()
	token := c.Request.Header.Get("Authorization")
	userID := libs.GetUserIDFromToken(token)
	if userID != "" {
		resultFind := db.Where("id = ?", userID).First(&userModel)
		if resultFind.Error == nil || errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
			if resultFind.RowsAffected <= 0 {
				status = http.StatusNotFound
				msg = "userid is not found"
			}
		} else {
			status = http.StatusInternalServerError
			msg = resultFind.Error.Error()
		}
	} else {
		status = http.StatusUnprocessableEntity
		msg = "userid is empty"
	}
	if status == 200 {
		now := time.Now()
		currentDate := now.Format("2006-01-02")
		var totalCount int64
		totalCount = 0
		result := db.Debug().Where("user_id = ? AND created_date = ?", userID, currentDate).Find(&taskModels).Count(&totalCount)
		if result.Error != nil {
			status = 500
			msg = result.Error.Error()
		} else {
			// @TODO if MaxToDo == 0 then add no limit
			if int(totalCount) < userModel.MaxToDo || userModel.MaxToDo <= 0 {
				body, _ := ioutil.ReadAll(c.Request.Body)
				var JSONObject map[string]interface{}
				json.Unmarshal([]byte(string(body)), &JSONObject)
				taskModel.PassBodyJSONToModel(JSONObject)
				taskModel.UserID = userID
				resultCreate := db.Create(&taskModel)
				if resultCreate.Error != nil {
					status = 500
					msg = resultCreate.Error.Error()
				}
			} else {
				status = http.StatusUnprocessableEntity
				msg = "maximum todo for today"
			}
		}
	}
	if status == 200 {
		msg = "Success"
		responseData = gin.H{
			"status": status,
			"data":   taskModel,
			"msg":    msg,
		}
	} else {
		if msg == "" {
			msg = "Error"
		}
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}
	libs.APIResponseData(c, status, responseData)
}
