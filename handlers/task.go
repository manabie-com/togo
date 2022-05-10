package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/utils"
)

const (
	COMPARE_DATE_FORMAT string = "2006-04-02"
)

type TaskHandler struct {
	models *models.Models
}

func NewTaskHandler(models *models.Models) *TaskHandler {
	return &TaskHandler{
		models: models,
	}
}

func (th *TaskHandler) UpdateUserTask(c *gin.Context) {
	var err error
	var user *models.User
	// Get request body data
	userTask := &models.UserTask{}
	if err = c.ShouldBind(userTask); err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			ErrorMsg:   err.Error(),
		})
		return
	}

	if user, err = th.models.User.GetUserByUserId(c.Request.Context(), userTask.UserID); err != nil {
		utils.ResponseJson(c, http.StatusInternalServerError, &utils.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Server internal error when getting UserID %s", userTask.UserID),
			ErrorMsg:   err.Error(),
		})
		return
	}

	if user != nil {
		/* Refill DailyTasksLimit value in next day */
		nowDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
		lastUpdated := time.Date(user.LastUpdatedTask.Year(), user.LastUpdatedTask.Month(), user.LastUpdatedTask.Day(), 0, 0, 0, 0, user.LastUpdatedTask.Location())
		fmt.Println("Now", nowDay)
		fmt.Println("Last", lastUpdated)
		if nowDay.After(lastUpdated) {
			fmt.Println("HERE")
			user.DailyTasksLimit = user.MaxDailyTasks
		}

		if user.DailyTasksLimit == 0 || nowDay.Before(lastUpdated) {
			utils.ResponseJson(c, http.StatusNotImplemented, &utils.ErrorResponse{
				StatusCode: http.StatusNotImplemented,
				Message:    fmt.Sprintf("Daily TODO tasks for UserID %s had been reached", user.UserID),
				ErrorMsg:   "",
			})
			return
		}
	} else {
		if userTask.MaxDailyLimit == 0 {
			userTask.MaxDailyLimit = models.DEFAULT_MAX_DAILY_TASKS
		}

		if user, err = th.models.User.CreateUser(c.Request.Context(), userTask.UserID, userTask.MaxDailyLimit); err != nil {
			utils.ResponseJson(c, http.StatusInternalServerError, &utils.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("Failed to create new UserID %s", user.UserID),
				ErrorMsg:   err.Error(),
			})
			return
		}
	}

	// Create task with decreasing DailyTasksLimit value
	if err = th.createTaskWithDecreaseLimit(c, user, userTask.TodoTask); err != nil {
		utils.ResponseJson(c, http.StatusInternalServerError, &utils.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Daily TODO tasks for UserID %s had been reached", user.UserID),
			ErrorMsg:   err.Error(),
		})
		return
	}

	utils.ResponseJson(c, http.StatusOK, true)
}

func (th *TaskHandler) createTaskWithDecreaseLimit(c *gin.Context, user *models.User, todoTask string) (err error) {
	// TODO task is valid to be added
	if _, err = th.models.Task.CreateTask(c.Request.Context(), user.UserID, todoTask); err != nil {
		utils.ResponseJson(c, http.StatusInternalServerError, &utils.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to create TODO task for UserID %s\n", user.UserID),
			ErrorMsg:   err.Error(),
		})
		return
	}

	// Decrease daily task limit and update time
	user.DailyTasksLimit--
	user.LastUpdatedTask = time.Now()

	// Update user with new DailyTaskLimit
	if _, err = th.models.User.UpdateUser(c.Request.Context(), user); err != nil {
		utils.ResponseJson(c, http.StatusInternalServerError, &utils.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to create TODO task for UserID %s\n", user.UserID),
			ErrorMsg:   err.Error(),
		})
		return
	}
	return nil
}
