package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/storages"
)

const TASKS_PER_DAY_CACHE_KEY = "tasksperday:%s:%s"
const LIMIT_PER_DAY = 5

type TaskService struct {
	Redis redis.Cmdable
	storages.Storage
}

func value(p string) sql.NullString {
	return sql.NullString{
		String: p,
		Valid:  true,
	}
}

func (s *TaskService) ListTasks(ctx context.Context, userID string, createdDate string) ([]*storages.Task, error) {
	return s.Storage.RetrieveTasks(
		ctx,
		value(userID),
		value(createdDate),
	)
}

func (s *TaskService) AddTask(ctx context.Context, body io.ReadCloser, userID string) (*storages.Task, *common.AppError) {
	today := time.Now().Format("2006-01-02")
	limitCacheKey := fmt.Sprintf(TASKS_PER_DAY_CACHE_KEY, today, userID)

	t := &storages.Task{}
	err := json.NewDecoder(body).Decode(t)
	if err != nil {
		return nil, &common.AppError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = today

	err = s.Storage.AddTask(ctx, t)
	if err == nil {
		s.Redis.IncrBy(ctx, limitCacheKey, 1)
	}

	return t, nil
}

func (s *TaskService) IsReachedLimit(ctx context.Context, userID string) *common.AppError {
	today := time.Now().Format("2006-01-02")
	limitCacheKey := fmt.Sprintf(TASKS_PER_DAY_CACHE_KEY, today, userID)

	val, err := s.Redis.Get(ctx, limitCacheKey).Result()
	if err != nil {
		return nil
	}

	count, err := strconv.Atoi(val)
	if err != nil {
		return &common.AppError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	if count >= LIMIT_PER_DAY {
		return &common.AppError{
			Code: http.StatusTooManyRequests,
			Err:  errors.New("daily tasks limit exceeded"),
		}
	}

	return nil
}
