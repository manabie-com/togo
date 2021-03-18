package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
	"time"
)

type userAuthKey int8

func (s *ToDoService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	var parDate string
	if parDate = req.FormValue("created_date"); len(parDate) == 0 {
		parDate = time.Now().Format("2006-01-02")
	}

	tasks, err := s.Store.RetrieveTasks(
		Value(id),
		Value(parDate),
	)

	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		writeHeader(&resp, http.StatusInternalServerError)
		json := make(map[string]string)
		json["error"] = err.Error()
		responseJson(&resp, json)

		return
	}

	json := make(map[string][]*storages.Task)
	json["data"] = tasks
	responseJson(&resp, json)
}

func (s *ToDoService) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}

	//err := json.NewDecoder(req.Body).Decode(t)
	//defer req.Body.Close()
	//if err != nil {
	//	writeHeader(&resp, http.StatusInternalServerError)
	//	return
	//}

	//userID, _ := userIDFromCtx(req.Context())
	userID := "firstUser"
	t.ID = uuid.New().String()
	t.UserID = userID
	t.Content = req.FormValue("content")
	now := time.Now()
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")
	isOutOfTasks := s.Store.CheckNumTasksInDay(Value(t.UserID), Value(t.CreatedDate))
	if isOutOfTasks {
		writeHeader(&resp, http.StatusMethodNotAllowed)
		json := make(map[string]string)
		json["error"] = "Reached max tasks to do in day"
		responseJson(&resp, json)

		return
	}

	err := s.Store.AddTask(t)
	if err != nil {
		writeHeader(&resp, http.StatusInternalServerError)
		json := make(map[string]string)
		json["error"] = err.Error()
		responseJson(&resp, json)

		return
	}

	json := make(map[string]*storages.Task)
	json["data"] = t
	responseJson(&resp, json)
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
