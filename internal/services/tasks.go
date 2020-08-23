package services

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/phuwn/togo/internal/storages"
	"github.com/phuwn/togo/util"
)

var (
	CreateTaskLimitErrMsg  = "pq: reach create_task_limit"
	CreateTaskLimitErrResp = "User has reached create_task limit of the day please try again tomorrow"
	UnknownErrResp         = "Unknown error occurs, please try again later"
	InvalidBodyErrResp     = "invalid body format"
)

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		if err == io.EOF {
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": InvalidBodyErrResp,
			})
			return
		}
		resp.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to add task, unknown error occurs", err)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = util.NewUUID()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		errMsg := err.Error()
		switch errMsg {
		case CreateTaskLimitErrMsg:
			resp.WriteHeader(http.StatusTooManyRequests)
			errMsg = CreateTaskLimitErrResp
		default:
			resp.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to add task, unknown error occurs", errMsg)
			errMsg = UnknownErrResp
		}

		json.NewEncoder(resp).Encode(map[string]string{
			"error": errMsg,
		})
		return
	}

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}
