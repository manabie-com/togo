package persister

import (
	"context"

	"github.com/trangmaiq/togo/internal/model"
)

type Persister interface {
	CreateTask(ctx context.Context, task *model.Task) error
}
