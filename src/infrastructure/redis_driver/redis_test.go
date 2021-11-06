package redis_driver

import (
	"os"
	"testing"

	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	redisDriver := RedisDriver{}
	err := redisDriver.Setup(RedisConfiguration{
		Addr: os.Getenv(settings.RedisHost) + ":" + os.Getenv(settings.RedisPort),
	})
	assert.NoError(t, err)
	pong, err := RedisClient.Ping().Result()
	assert.NoError(t, err)
	assert.EqualValues(t, "PONG", pong)
}
