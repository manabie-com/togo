package redix

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"togo/internal/entity"
)

func (r *redisStore) GetTask(ctx context.Context, id int32, userId int32) (*entity.Task, error) {
	key := fmt.Sprintf("tasks/%d/%d", userId, id)

	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	task := entity.Task{}

	json.Unmarshal([]byte(result), &task)

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

func (r *redisStore) DeleteTask(ctx context.Context, task *entity.Task) error {
	key := fmt.Sprintf("tasks/%d/%d", task.UserID, task.ID)

	result, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
