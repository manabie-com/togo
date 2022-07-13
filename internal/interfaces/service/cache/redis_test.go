package cache

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisSuccess(t *testing.T) {
	location := DefaultRedisLocation
	options, err := redis.ParseURL(location)
	assert.NoError(t, err)

	options.Username = ""
	client := redis.NewClient(options)
	ctx := context.Background()

	_, err = client.Set(ctx, UserQuotaKey(1), 3, DefaultUserQuotaTTL).Result()
	assert.NoError(t, err)

	quota, err := client.Get(ctx, UserQuotaKey(1)).Int()
	assert.NoError(t, err)
	assert.Equal(t, 3, quota)
}
