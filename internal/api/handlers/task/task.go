package task

import (
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/internal/api/handlers"
	"example.com/m/v2/internal/models"
	"example.com/m/v2/internal/pkg/responses"
	"example.com/m/v2/internal/repositories/task"

	"github.com/gin-gonic/gin"
)

func AddTask(service handlers.MainUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userInfo, err := handlers.GetUserInfoFromToken(ctx)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusUnauthorized, "Access denied")
			return
		}

		// Get task per day of user
		maxTaskOfUser, err := strconv.Atoi(userInfo.MaxTaskPerDay)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusInternalServerError, "Fail Parse String to Int")
			return
		}
		createDate := time.Now().Format("2006-01-02")

		// Find all tasks of user with userID and createDate
		tasks, err := service.Task.FindTaskByUser(userInfo.ID, createDate)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusBadRequest, "Fail FindTaskByUser")
			return
		}

		// Check number of task
		if IsMaxTaskPerDay(len(tasks), maxTaskOfUser) {
			responses.ResponseForError(ctx, nil, http.StatusInternalServerError, "Number tasks of User is maximum")
			return
		}

		// Mapping Request
		input := models.Task{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			responses.ResponseForError(ctx, err, http.StatusBadRequest, "Fail BindJSON user")
			return
		}
		input.UserID = userInfo.ID
		input.CreateDate = createDate

		// New
		out := task.NewTask(input)

		// Create Task
		if err := service.Task.AddTask(out); err != nil {
			responses.ResponseForError(ctx, err, http.StatusBadRequest, "Fail Add Task")
			return
		}

		responses.ResponseForOK(ctx, http.StatusOK, nil, "Success")
	}
}

func IsMaxTaskPerDay(numTaskPerDay int, maxTaskOfUser int) bool {
	return numTaskPerDay >= maxTaskOfUser
}