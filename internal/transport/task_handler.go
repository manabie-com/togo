package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase"
)

const (
	LAYOUT = "2006-01-02"
)

type TaskHandler struct {
	TUsecase usecase.TaskUsecase
}

func NewTaskHandler(us usecase.TaskUsecase) TaskHandler {
	return TaskHandler{
		TUsecase: us,
	}
}

func (t *TaskHandler) ListTasks(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user_id, _ := userIDFromCtx(ctx)
	created_date := req.FormValue("created_date")
	created_dateT, _ := time.Parse(LAYOUT, created_date)
	tasks, err := t.TUsecase.ListTasks(ctx, user_id, created_dateT)

	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(resp).Encode(map[string][]storages.Task{
		"data": tasks,
	})
}

func (t *TaskHandler) AddTask(resp http.ResponseWriter, req *http.Request) {
	var task storages.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := userIDFromCtx(req.Context())
	task.UserID = userID
	task.CreatedDate = time.Now()
	err = t.TUsecase.AddTask(req.Context(), &task)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": &task,
	})
}
func (t *TaskHandler) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("user_id")
	pass := req.FormValue("password")

	val, err := t.TUsecase.ValidateUser(req.Context(), id, pass)
	if err != nil || val {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := CreateToken(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
