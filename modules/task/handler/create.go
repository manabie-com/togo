package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/models"
	"github.com/khoale193/togo/modules/task/service"
	"github.com/khoale193/togo/pkg/app"
	"github.com/khoale193/togo/pkg/e"
	"github.com/khoale193/togo/pkg/util"
)

type TaskForm struct {
	Name string `json:"name" validate:"required"`
}

func CreateTask(c *gin.Context) {
	var appG = app.Gin{C: c}
	form := CreateTaskFormValidator{}
	if httpCode, err := form.BindAndValid(c); err != nil {
		appG.ResponseError(httpCode, app.NewValidatorError(err), nil)
		return
	}
	if err := form.service.CreateTask(); err != nil {
		appG.Response(c, http.StatusBadRequest, e.Msg[e.ERROR], e.ERROR, nil)
		return
	}
	appG.Response(c, http.StatusOK, e.GetMsg(e.SUCCESS), e.SUCCESS, nil)
}

type CreateTaskFormValidator struct {
	TaskForm
	service service.Task
}

func (v *CreateTaskFormValidator) BindAndValid(c *gin.Context) (int, error) {
	err := app.BindAndValid(c, &v.TaskForm)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if (&models.Member{ID: int64(GetCurrentUserID(c))}).IsExceedLimitTaskPerDay() {
		return http.StatusBadRequest, errors.New(e.Msg[e.ERROR_IN_EXCEED_LIMIT_TASK_ADD_PER_DAY])
	}
	v.service.Name = v.Name
	v.service.MemberID = GetCurrentUserID(c)
	return http.StatusOK, nil
}

func GetCurrentUserID(c *gin.Context) int {
	return util.ExtractMemberID(c.Request.Header.Get(e.UserAuth))
}
