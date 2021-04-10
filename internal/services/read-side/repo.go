package read_side

import (
	"context"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
)

type ReadRepo interface {
	GetTaskList(ctx context.Context, values url.Values) ([]*TaskList, string, error)
}

type TaskList struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
