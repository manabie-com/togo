package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/daos"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/utils"
)

func ValidateDailyTaskLimit(accountID uuid.UUID, maxDailyTasksCount uint) bool {
	startTime, endTime, _ := utils.GetTimesByPeriod("daily")
	taskDAO := daos.GetTaskDAO()
	taskCount, _ := taskDAO.CountTaskByAccountIDAndPeriod(accountID, startTime, endTime)
	return taskCount < maxDailyTasksCount
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var (
		message = "OK"
	)
	//parse from token
	ctx := r.Context()
	accountID, _ := uuid.Parse(fmt.Sprint(ctx.Value("account_id")))
	maxDailyTasksCount, _ := strconv.ParseUint(fmt.Sprint(r.Context().Value("max_daily_tasks_count")), 10, 0)
	//validate whether the user has reached the daily task limit
	isValidLimit := ValidateDailyTaskLimit(accountID, uint(maxDailyTasksCount))
	if !isValidLimit {
		config.ResponseWithError(w, "You have reached the daily task limit", nil)
		return
	}
	//create task
	task := models.Task{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	task.AccountID = accountID
	taskDAO := daos.GetTaskDAO()
	result, err := taskDAO.CreateTask(task)
	if err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	config.ResponseWithSuccess(w, message, result)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var (
		// params   = mux.Vars(r)
		message = "OK"
	)
	startTime, endTime, _ := utils.GetTimesByPeriod("daily")
	var err error
	if len(r.URL.Query()["start"]) > 0 {
		startTime, err = time.Parse("02-01-2006", r.URL.Query()["start"][0])
		if err != nil {
			config.ResponseWithError(w, "invalid start date (dd-mm-yyyy)", err)
			return
		}
		//if both start and end time are inputted
		if len(r.URL.Query()["end"]) > 0 {
			endTime, err = time.Parse("02-01-2006", r.URL.Query()["end"][0])
			if err != nil {
				config.ResponseWithError(w, "invalid end date (dd-mm-yyyy)", err)
				return
			}
		} else {
			year := startTime.Year()
			month := startTime.Month()
			day := startTime.Day()
			location := startTime.Location()
			endTime = time.Date(year, month, day, 23, 59, 59, 0, location)
		}
	} else if len(r.URL.Query()["end"]) > 0 {
		endTime, _ = time.Parse("02-01-2006", r.URL.Query()["end"][0])
		if err != nil {
			config.ResponseWithError(w, "invalid end date (dd-mm-yyyy)", err)
			return
		}
	}
	taskDAO := daos.GetTaskDAO()
	tasks, err := taskDAO.GetTasksByPeriod(startTime, endTime)
	if err != nil {
		config.ResponseWithError(w, "Get tasks failed", err)
		return
	} else {
		config.ResponseWithSuccess(w, message, tasks)
	}
}
