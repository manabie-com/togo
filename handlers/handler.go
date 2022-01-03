package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/namnhatdoan/togo/constants"
	"github.com/namnhatdoan/togo/db"
	"github.com/namnhatdoan/togo/settings"
	"github.com/namnhatdoan/togo/utils"
	"strings"
	"time"
)

func CreateTask(c *gin.Context) {
	// Binding request
	req := CreateTaskRequest{}
	if success := bindingBodyData(c, &req); !success {
		return
	}

	// Validate data
	if err := normalizeAndValidateCreateTaskReq(&req); err != nil {
		badRequestError(c, err.Error())
		return
	}

	// Get today tasks
	if task, err := db.AddNewTask(req.Email, req.Task); err != nil {
		serverError(c, err.Error())
	} else {
		responseSuccess(c, task)
	}
}

func normalizeAndValidateCreateTaskReq(req *CreateTaskRequest) error {
	// Trim email and task
	req.Email = strings.TrimSpace(req.Email)
	req.Task = strings.TrimSpace(req.Task)

	// Validate
	if err := validateEmail(req.Email); err != nil {
		return errors.New(constants.InvalidEmail)
	}
	if err := validateTask(req.Task); err != nil {
		return err
	}

	return nil
}

func SetConfig(c *gin.Context) {
	// Binding request
	req := SetConfigRequest{}
	if success := bindingBodyData(c, &req); !success {
		return
	}

	// Validate data
	if err := normalizeAndValidateSetConfigReq(&req); err != nil {
		badRequestError(c, err.Error())
		return
	}

	// Get today tasks
	if conf, err := db.SetConfig(req.Email, req.Limit, req.date); err != nil {
		serverError(c, err.Error())
	} else {
		responseSuccess(c, conf)
	}
}

func normalizeAndValidateSetConfigReq(req *SetConfigRequest) error {
	// Trim email and task
	req.Email = strings.TrimSpace(req.Email)

	// Validate
	if err := validateEmail(req.Email); err != nil {
		return errors.New(constants.InvalidEmail)
	}
	if req.Limit == 0 {
		return errors.New(constants.MissingTaskLimit)
	}
	if req.Limit > constants.MaxTaskPerDay{
		return errors.New(constants.LimitTaskOverMaxValue)
	}
	if req.Date != "" {
		// Parse inputted date
		date, err := time.ParseInLocation(constants.DateStringFormat, req.Date, time.Local)
		if err != nil {
			settings.GetLogger().WithField("date", req.Date).WithError(err).Error("Parse date fail")
			return errors.New(constants.InvalidDate)
		}
		req.date = date
	} else {
		// Set req date to today
		req.date = utils.GetCurrentDate()
	}

	return nil
}
