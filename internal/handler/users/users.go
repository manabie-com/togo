package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	r_users "github.com/manabie-com/togo/internal/reqres/users"
	u_users "github.com/manabie-com/togo/internal/usecase/users"
	"github.com/manabie-com/togo/pkg/exception"
	"github.com/manabie-com/togo/pkg/validate"
)

type UserHandler struct{}

func (u *UserHandler) AssignTasks(c *gin.Context) {
	// get path parameter
	paramUserID := c.Param("id")
	userID, err := strconv.ParseInt(paramUserID, 10, 16)
	if err != nil {
		exception.ThrownError(c, err)
		return
	}

	// parse request body
	request := &r_users.AssignTaskRequest{
		UserID: int16(userID),
	}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		exception.ThrownError(c, err)
		return
	}

	// validate
	err = validate.Struct(request)
	if err != nil {
		exception.ThrownError(c, err)
		return
	}

	// usecase
	err = u_users.AssignUserTasks(int16(userID), request)
	if err != nil {
		exception.ThrownError(c, err)
		return
	}

	c.String(http.StatusOK, "")
}
