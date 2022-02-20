package redisrepo

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"togo/internal/domain"
	"togo/internal/repository"

	"github.com/go-redis/redis/v8"
)

type taskLimitRepository struct {
	rdb    *redis.Client
	prefix string
}

// NewTaskLimitRepository repository constructor
func NewTaskLimitRepository(rdb *redis.Client, prefix string) repository.TaskLimitRepository {
	return &taskLimitRepository{
		rdb,
		prefix,
	}
}

func (r taskLimitRepository) Increase(ctx context.Context, userID uint, limit int) (int, error) {
	count := 0
	t := time.Now().UTC()
	// Last of current month
	exp := t.AddDate(0, 1, -t.Day())
	// Day-specific user task limit key
	key := fmt.Sprintf("%s:%v:%v", r.prefix, userID, t.Day())
	// Check the current tasks submitted
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, fmt.Errorf("taskLimitRepository:Increase: %w", err)
	}
	if val != "" {
		if count, err = strconv.Atoi(val); err != nil {
			return 0, fmt.Errorf("taskLimitRepository:Increase: %w", err)
		}
	}
	if count >= limit {
		return count, domain.ErrTaskLimitExceed
	}
	// Add task count transaction
	pipe := r.rdb.TxPipeline()
	incrOp := pipe.Incr(ctx, key)
	pipe.ExpireAt(ctx, key, exp)
	if _, err := pipe.Exec(ctx); err != nil {
		return 0, fmt.Errorf("taskLimitRepository:Increase: %w", err)
	}
	return int(incrOp.Val()), nil
}
