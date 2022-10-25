package api

import (
	"net/http"
	"togo/msg"
	"togo/services"
	"togo/utils"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email" valid:"email~Email is not valid"`
	Password string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
}

func Login(c *gin.Context) {
	appG := Gin{C: c}
	var loginReq LoginReq
	isValid := appG.BindAndValidate(&loginReq)
	if isValid {
		service := services.UserReq{Email: loginReq.Email, Password: loginReq.Password}
		user, err := service.Login()
		if err != nil {
			appG.Response(http.StatusUnauthorized, false, msg.GetMsg(msg.ERROR_AUTH_FAIL), nil, nil)
			return
		}
		j := utils.NewJWT()
		tokenInfo, err := j.GenerateToken(user.ID, user.Email, user.FullName)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		user.Token = tokenInfo.Token
		user.ExpiredAt = tokenInfo.ExpiredAt
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), user, nil)
	}
}
