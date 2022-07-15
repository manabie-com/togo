package cache

import (
	"context"
	"strconv"

	"github.com/datshiro/togo-manabie/internal/interfaces/consts"
	"github.com/go-redis/redis/v9"
)

func (r Redis) SetUserQuota(ctx context.Context, userID int, taskCount int) error {
	_, err := r.client.Set(ctx, UserQuotaKey(userID), taskCount, DefaultUserQuotaTTL).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) GetUserQuota(ctx context.Context, userID int) (int, error) {
	val, err := r.client.Get(ctx, UserQuotaKey(userID)).Result()
	if err != nil {
		if err == redis.Nil { // Not found
			return 0, nil
		}
		return 0, err
	}
	return strconv.Atoi(val)
}
func UserQuotaKey(userID int) string {
	return consts.UserTaskQuotaKey + strconv.Itoa(userID)
}
