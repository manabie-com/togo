package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/manabie-com/togo/internal/storages"
)

type taskCountStore struct {
	client *redis.Client
}

func getKey(userID, date string) string {
	return fmt.Sprintf("%s$|$%s", userID, date)
}

func getTimeRemain() time.Duration {
	t := time.Now()
	tA, _ := time.Parse("2006-01-02", fmt.Sprintf("%v-%v-%v", t.Year(), t.Month(), t.Day()))
	return tA.Add(24 * time.Hour).Sub(tA)

}

func (s *taskCountStore) CreateIfNotExists(ctx context.Context, userID, date string) error {
	key := getKey(userID, date)

	cmd := s.client.SetNX(ctx, key, 0, getTimeRemain())
	return cmd.Err()
}
func (s *taskCountStore) Inc(ctx context.Context, userID, date string) (int, error) {
	key := getKey(userID, date)
	cmd := s.client.Incr(ctx, key)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	return int(cmd.Val()), nil

}

func (s *taskCountStore) Desc(ctx context.Context, userID, date string) error {
	key := getKey(userID, date)
	cmd := s.client.Decr(ctx, key)
	return cmd.Err()

}
func NewTaskCountStore(client *redis.Client) storages.TaskCountStore {
	return &taskCountStore{
		client: client,
	}
}
