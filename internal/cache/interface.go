package cache

import (
	"context"
	"time"
)

type Cache interface {
	AddToken(ctx context.Context, token string, exp time.Duration) error
	ValidateToken(ctx context.Context, token string) error
}
