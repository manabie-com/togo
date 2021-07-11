package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/api/task/storages"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Store storages.Store
}

func (s *Task) List(ctx context.Context, userID, createdDate string, page, limit int) ([]*storages.Task, error) {
	tasks, err := s.Store.RetrieveTasks(ctx, userID, createdDate, page, limit)
	if err != nil {
		logger.MBErrorf(ctx, "task storage failed to retrieve tasks of user_id %s created_date %s: %v", userID, createdDate, err)
		return nil, errors.New(dictionary.FailedGetRetrieveTasks)
	}

	return tasks, nil
}

func (s *Task) Add(resp http.ResponseWriter, req *http.Request) {
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

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
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
