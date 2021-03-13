package interfaces

import (
	"context"
	"time"

	"github.com/valonekowd/togo/domain/entity"
)

type DataSource struct {
	User UserDataSource
	Task TaskDataSource
}

type UserDataSource interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Add(ctx context.Context, u *entity.User) error
}

type TaskDataSource interface {
	List(ctx context.Context, userID string, createdDate time.Time) ([]*entity.Task, error)
	Add(ctx context.Context, t *entity.Task) error
}
