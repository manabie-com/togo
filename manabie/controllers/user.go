package controllers

import (
	"errors"
	"manabie/manabie/databases/drivers/sqlite"
	libs "manabie/manabie/helpers"
	"manabie/manabie/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	jwt "github.com/dgrijalva/jwt-go"
)

func Login(c *gin.Context) {
	defer libs.RecoverError(c)
	var (
		status           = 200
		msg              string
		responseData     = gin.H{}
		token            string
		userID, password string
		userModel        models.User
		err              error
	)
	userID, _ = libs.GetQueryParam("user_id", c)
	password, _ = libs.GetQueryParam("password", c)
	db := sqlite.Connect()
	resultFind := db.Where("id = ? AND password = ?", userID, password).First(&userModel)
	if resultFind.Error == nil || errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
		if resultFind.RowsAffected <= 0 {
			status = http.StatusUnauthorized
			msg = "incorrect user_id/pwd"
		} else {
			atClaims := jwt.MapClaims{}
			atClaims["user_id"] = userID
			atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			token, err = at.SignedString([]byte(os.Getenv("JWT_TOKEN")))
			if err != nil {
				status = 500
				msg = err.Error()
			}
		}
	} else {
		status = 500
	}
	if status == 200 {
		msg = "Success"
		responseData = gin.H{
			"status": status,
			"data":   token,
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
