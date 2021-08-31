package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/repo"
)

// TaskService implement
type TaskService struct {
	TaskStore *repo.TaskStore
	UserStore *repo.UserStore
}

func (this *TaskService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := this.TaskStore.RetrieveTasks(req.Context(), id, value(req, "created_date"))

	resp.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

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

func (this *TaskService) CreateTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	isExceeds, err := this.isTaskExceeds(req.Context())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	if isExceeds {
		resp.WriteHeader(http.StatusUnavailableForLegalReasons)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Task exceeds",
		})
		return
	}

	err = this.TaskStore.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}

func (this *TaskService) isTaskExceeds(ctx context.Context) (bool, error) {
	userID, _ := userIDFromCtx(ctx)
	count, err := this.TaskStore.CountTask(ctx, userID, time.Now().Format("2006-01-02"))
	if err != nil {
		return true, err
	}

	user, err := this.UserStore.RetrieveUser(ctx, userID)
	if err != nil {
		return true, err
	}
	if count >= user.MaxTodo {
		return true, nil
	}
	return false, nil
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
