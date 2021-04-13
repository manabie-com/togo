package transport

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/entities"
	"net/http"
	"time"
)

type UserAuthKey int8
type Transport struct{}

func NewTransport() *Transport {
	return &Transport{}
}
func (t *Transport) GetValue(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (t *Transport) GetToken(req *http.Request) string {
	return req.Header.Get("Authorization")
}

func (t *Transport) GetUserIDFromCtx(ctx context.Context) string {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	if ok {
		return id
	}
	return ""
}

func (t *Transport) GetTask(req *http.Request) (*entities.Task, error) {
	task := &entities.Task{}
	err := json.NewDecoder(req.Body).Decode(task)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	userID := t.GetUserIDFromCtx(req.Context())
	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = now.Format("2006-01-02")
	return task, nil
}
