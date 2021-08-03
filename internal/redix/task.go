package redix

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"togo/internal/entity"
)

func TaskKey(id int32) string {
	return fmt.Sprintf("tasks/%d", id)
}

func (r *redisStore) GetTask(ctx context.Context, id int32) (*entity.Task, error) {
	key := TaskKey(id)

	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	task := entity.Task{}

	_ = json.Unmarshal([]byte(result), &task)

	return &task, nil
}

func (r *redisStore) SetTask(ctx context.Context, task *entity.Task) error {
	key := fmt.Sprintf("tasks/%d/%d", task.UserID, task.ID)

	out, err := json.Marshal(task)
	if err != nil {
		return err
	}

	cmd := r.client.Set(ctx, key, out, time.Minute*10)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *redisStore) DeleteTask(ctx context.Context, id int32) error {
	key := TaskKey(id)

	if _, err := r.client.Del(ctx, key).Result(); err != nil {
		return err
	}

	return nil
}
