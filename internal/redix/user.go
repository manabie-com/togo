package redix

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"togo/internal/entity"
)

func (r *redisStore) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	key := fmt.Sprintf("users/%s", username)

	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user := entity.User{}

	_ = json.Unmarshal([]byte(result), &user)

	return &user, nil
}

func (r *redisStore) GetUser(ctx context.Context, id int32) (*entity.User, error) {
	key := fmt.Sprintf("users/%d", id)

	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	u := entity.User{}

	_ = json.Unmarshal([]byte(result), &u)

	return &u, nil
}

func (r *redisStore) SetUser(ctx context.Context, user *entity.User) error {
	out, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("users/%d", user.ID)
	cmd := r.client.Set(ctx, key, out, time.Minute*10)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	key = fmt.Sprintf("users/%s", user.Username)
	cmd = r.client.Set(ctx, key, out, time.Minute*10)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
