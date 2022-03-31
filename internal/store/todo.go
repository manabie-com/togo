package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vchitai/togo/internal/models"
	"gorm.io/gorm"
)

type ToDo interface {
	Record(ctx context.Context, list []*models.ToDo) error
	GetConfig(ctx context.Context, userID string) (*models.ToDoConfig, error)
	GetUsedCount(ctx context.Context, userID string, day time.Time) (int64, error)
	IncreaseUsedCount(ctx context.Context, userID string, day time.Time, by int64) (int64, error)
	DecreaseUsedCount(ctx context.Context, userID string, day time.Time, by int64) (int64, error)
}

var _ ToDo = &toDo{}

func NewToDo(db *gorm.DB, redisCli *redis.Client) *toDo {
	return &toDo{db: db, redisCli: redisCli}
}

type toDo struct {
	db       *gorm.DB
	redisCli *redis.Client
}

func (t *toDo) IncreaseUsedCount(ctx context.Context, userID string, day time.Time, by int64) (int64, error) {
	var key = BuildUserDailyUsedCount(userID, day)
	res, err := t.redisCli.IncrBy(ctx, key, by).Result()
	if err != nil {
		return 0, err
	}
	// if first time, set expire
	if res == by {
		_ = t.redisCli.Expire(ctx, key, dailyUsedCountTTL)
	}
	return res, nil
}

func (t *toDo) DecreaseUsedCount(ctx context.Context, userID string, day time.Time, by int64) (int64, error) {
	var key = BuildUserDailyUsedCount(userID, day)
	res, err := t.redisCli.DecrBy(ctx, key, by).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (t *toDo) GetConfig(ctx context.Context, userID string) (*models.ToDoConfig, error) {
	var res models.ToDoConfig
	var err = t.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		First(&res).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &res, err
}

func (t *toDo) GetUsedCount(ctx context.Context, userID string, day time.Time) (int64, error) {
	var key = BuildUserDailyUsedCount(userID, day)
	res, err := t.redisCli.Get(ctx, key).Int64()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (t *toDo) Record(ctx context.Context, list []*models.ToDo) error {
	return t.db.WithContext(ctx).Save(list).Error
}

func BuildUserDailyUsedCount(userID string, day time.Time) string {
	return fmt.Sprintf("daily-used-count:%s:%d", userID, day.Unix())
}

const dailyUsedCountTTL = 24 * time.Hour
