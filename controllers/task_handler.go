package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	_const "togo-thdung002/const"
	"togo-thdung002/entities"
	"togo-thdung002/entities/response"
	"togo-thdung002/utils"
)

func apiGetTask(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		userID := c.QueryParam("userid")
		id, err := strconv.Atoi(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}
		tasks, err := s.db.GetTaskFilterBy(uint(id), "")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}

		return c.JSON(http.StatusOK, resp.Success(tasks))
	})
}

func apiPostTask(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err))
		}
		var task entities.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}
		if err := task.Validate(); err != nil {
			log.Error("Error on validate struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(_const.ErrValidate))
		}
		if task.UserID == 0 {
			log.Error("User not found", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(_const.ErrUserNotFound))
		}
		if ok := utils.IsDateValue(task.Date); !ok {
			log.Error("Date type is not right - make sure it is dd-mm-yyyy")
			return c.JSON(http.StatusInternalServerError, resp.Error(_const.ErrDateType))
		}
		tasks, err := s.db.GetTaskFilterBy(task.UserID, task.Date)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}
		user, err := s.db.GetUser(task.UserID)
		if len(tasks) > user.Limit-1 {
			return c.JSON(http.StatusInternalServerError, resp.Error(_const.ErrLimitedTask))
		}

		createdID, err := s.db.CreateTask(&task)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}

		return c.JSON(http.StatusOK, resp.Success(createdID))
	})
}
