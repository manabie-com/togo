package redix

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
	if err = r.client.Set(ctx, key, out, time.Minute*10).Err(); err != nil {
		return err
	}

	key = fmt.Sprintf("users/%s", user.Username)
	if err = r.client.Set(ctx, key, out, time.Minute*10).Err(); err != nil {
		return err
	}

	return nil
}

func KeyCount(userID int32) string {
	t := time.Now()

	return fmt.Sprintf("user_count/%d/%s", userID, t.Format("2006-01-02"))
}

func (r *redisStore) GetCountTaskToday(ctx context.Context, userID int32) (int32, error) {
	key := KeyCount(userID)

	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	number, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(number), nil
}

func (r *redisStore) SetCountTaskToday(ctx context.Context, userID int32, count int32) error {
	key := KeyCount(userID)

	if err := r.client.Set(ctx, key, count, time.Hour*24).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisStore) IncrCountTaskToday(ctx context.Context, userID int32) error {
	key := KeyCount(userID)

	if err := r.client.Incr(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisStore) DecrCountTaskToday(ctx context.Context, userID int32) error {
	key := KeyCount(userID)

	if err := r.client.Decr(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
