package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
)

type cacheRepository struct {
	getClient func(ctx context.Context) *redis.Client
	getDB     func(ctx context.Context) *gorm.DB
	totalTodo string
}

func NewCacheRepository(
	getClient func(ctx context.Context) *redis.Client,
	getDB func(ctx context.Context) *gorm.DB,
) CacheRepository {
	return &cacheRepository{
		getClient: getClient,
		getDB:     getDB,
		totalTodo: "total_todo",
	}
}

func (cr cacheRepository) GetTotalTodoByUserID(ctx context.Context, userID int64) (int, error) {
	total, err := cr.getClient(ctx).Get(ctx, fmt.Sprintf("%v:%v", cr.totalTodo, userID)).Int64()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, errors.Wrap(err, "get total todo")
	}

	if err != nil {
		err = cr.getDB(ctx).
			Unscoped().
			Model(&model.Todo{}).
			Where("user_id = ?", userID).
			Where("DATE(created_at) = ?", time.Now().Format("2006-01-02")).
			Count(&total).
			Error
		if err != nil {
			return 0, errors.Wrap(err, "get todo")
		}

		err = cr.SetTotalTodoByUserID(ctx, userID, int(total))
		if err != nil {
			return 0, err
		}
	}

	return int(total), nil
}

func (cr *cacheRepository) SetTotalTodoByUserID(ctx context.Context, userID int64, totalRequest int) error {
	err := cr.getClient(ctx).Set(ctx, fmt.Sprintf("%v:%v", cr.totalTodo, userID), totalRequest, -1).Err()
	if err != nil {
		return errors.Wrap(err, "set total todo")
	}

	return nil
}

func (cr *cacheRepository) ResetTotalTodo(ctx context.Context) error {
	rdb := cr.getClient(ctx)
	keys := rdb.Keys(ctx, fmt.Sprintf("%v:*", cr.totalTodo)).Val()

	for _, key := range keys {
		err := rdb.Set(ctx, key, 0, -1).Err()
		if err != nil {
			return errors.Wrap(err, "reset total todo")
		}
	}

	return nil
}
