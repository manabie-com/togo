package transport

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/manabie-com/togo/internal/pkg/timestamp"
	"github.com/manabie-com/togo/internal/storages"

	"github.com/google/uuid"
)

type LoginRequest struct {
	ID       string `form:"user_id"`
	Password string `form:"password"`
}

type CreateTaskRequest struct {
	Content string `json:"content"`
}

func parseCreateTaskParams(req *http.Request) (task storages.Task, err error) {
	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&task); err != nil {
		return task, err
	}

	userID, ok := userIDFromCtx(req.Context())
	if !ok {
		return task, errors.New("user_id is empty")
	}

	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = timestamp.GetCurrentTime()

	return
}
