package user

import (
	"github.com/gin-gonic/gin"
	jwt_service "manabie-com/togo/config/jwt-service"
	"manabie-com/togo/form"
	"manabie-com/togo/query"
	"manabie-com/togo/util"
	"net/http"
)

func LogIn(c *gin.Context) {
	var jwtLogInForm form.LogIn
	if err := c.ShouldBindJSON(&jwtLogInForm); err != nil {
		util.AbortJSONBadRequest(c)
		return
	}

	var user, err = query.UserByID(jwtLogInForm.Id)

	if err != nil {
		util.AbortUnauthorized(c, util.ERR_CODE_AUTH_WRONG_USER_NAME_OR_PASSWORD, "Wrong username or password!")
		return
	}

	if util.CompareHashPasswordAndPassword(user.Password, jwtLogInForm.Id, jwtLogInForm.Password) {
		util.AbortUnauthorized(c, util.ERR_CODE_AUTH_WRONG_USER_NAME_OR_PASSWORD, "Wrong username or password!")
		return
	}

	claim := jwt_service.Claims{
		Id:      jwtLogInForm.Id,
		MaxTodo: user.MaxTodo,
	}

	var token = jwt_service.CreateJwt(claim)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
