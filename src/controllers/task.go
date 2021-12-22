package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-api/service"
	"todo-api/src/helpers"
	"todo-api/src/models"

	"github.com/labstack/echo/v4"
)

func GetTasks(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("UserID").(string)

		var tasks []*models.Task
		err := s.DB.Where("user_id = ? and status = ?", userID, models.ActiveTaskStatus).
			Find(&tasks).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.BaseResponse{
			Data: tasks,
		})
	}
}

func CreateTask(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.Task
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{
				Message: err.Error(),
			})
		}

		userID := c.Get("UserID").(string)
		req.UserID = userID

		now := time.Now()
		var count int64
		err := s.DB.Model(models.Task{}).
			Where("user_id = ? and status = ?", userID, models.ActiveTaskStatus).
			Where("created_at >= ? and created_at < ?", helpers.BeginOfDate(now), helpers.EndOfDate(now)).
			Debug().
			Count(&count).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Message: err.Error(),
			})
		}

		if count+1 >= int64(s.Config.LimitTask) {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{
				Message: "Reached limit task per day",
			})
		}
		req.Status = models.ActiveTaskStatus
		err = s.DB.Model(models.Task{}).Create(&req).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(200, models.BaseResponse{
			Data: req,
		})
	}
}

func DeleteTask(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("UserID").(string)
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{
				Message: "Invalid id",
			})
		}

		affected := s.DB.Model(&models.Task{ID: id}).Where("user_id = ? and status = ?", userID, models.ActiveTaskStatus).
			Update("status", models.DeletedTaskStatus).RowsAffected
		if affected == 0 {
			return c.JSON(http.StatusUnprocessableEntity, models.BaseResponse{
				Message: "Delete task unsuccessful",
			})
		}

		return c.JSON(200, models.BaseResponse{
			Message: fmt.Sprintf("Deleted %d", id),
		})
	}
}
