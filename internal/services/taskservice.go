package services

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/middleware"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

type taskRequest struct {
	UserId      string
	CreatedDate sql.NullString `json:"created_date" binding:"required,alphanum"`
}

type taskListReponse struct {
	Data []storages.Task `json:"data"`
}

func (sc *ServiceController) taskListHandler(resp http.ResponseWriter, req *http.Request) {
	id, _ := middleware.UserIdFromCtx(req.Context())

	tasks, err := sc.Store.RetrieveTasks(req.Context(), postgres.RetrieveTasksParams{
		UserID:      sql.NullString{String: id, Valid: true},
		CreatedDate: value(req, "created_date"),
	})

	if err != nil {
		common.CommonResponse(resp, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	common.CommonResponse(resp, http.StatusOK, map[string][]storages.Task{
		"data": tasks,
	})
}

func (sc *ServiceController) updateTasksHandler(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		common.CommonResponse(resp, http.StatusBadRequest, map[string]string{
			"error": "bad_request",
		})
		return
	}

	now := time.Now()
	userID, _ := middleware.UserIdFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	countT, err := sc.Store.CountTaskPerDay(req.Context(), postgres.CountTaskPerDayParams{
		UserID:      t.UserID,
		CreatedDate: t.CreatedDate,
	})

	if err != nil {
		common.CommonResponse(resp, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if countT >= sc.Config.NLimit {
		common.CommonResponse(resp, http.StatusBadRequest, map[string]string{
			"error": "bad_request",
		})
		return
	}

	err = sc.Store.AddTask(req.Context(), t)

	if err != nil {
		common.CommonResponse(resp, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	common.CommonResponse(resp, http.StatusOK, map[string]storages.Task{
		"data": *t,
	})
}
