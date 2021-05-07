package user

import (
	"github.com/gin-gonic/gin"
	"manabie-com/togo/entity"
	"manabie-com/togo/query"
	"manabie-com/togo/util"
	"net/http"
)

func CreateOne(c *gin.Context) {
	var userForm entity.User
	if err := c.ShouldBindJSON(&userForm); err != nil {
		util.AbortJSONBadRequest(c)
		return
	}

	var _, err = query.UserByID(userForm.ID)

	if err == nil {
		util.AbortAlreadyExists(c, util.ERR_CODE_USER_EXISTED, "user is existed!")
		return
	}

	user := entity.User{
		ID:       userForm.ID,
		Password: util.HashPassword(userForm.ID, userForm.Password),
	}

	if userForm.MaxTodo <= 0 {
		user.MaxTodo = entity.DefaultMaxTodo
	} else {
		user.MaxTodo = userForm.MaxTodo
	}

	err = user.Create()

	user.Password = ""

	if err != nil {
		util.AbortUnexpected(c, util.ERR_CODE_DB_ISSUE, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
