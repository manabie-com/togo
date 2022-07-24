package handler

import (
	"context"
	"time"

	"github.com/trangmaiq/togo/internal/model"
)

//go:generate mockgen -source handler.go -destination handler_mock.go -package handler
type (
	Persister interface {
		CreateTask(ctx context.Context, task *model.Task) error
	}
	RateLimiter interface {
		AllowN(t time.Time, id string, n int64) bool
	}
	Dependencies interface {
		Persister() Persister
		RateLimiter() RateLimiter
	}
	Handler struct {
		d Dependencies
	}
)

func New(d Dependencies) *Handler {
	return &Handler{d: d}
}
