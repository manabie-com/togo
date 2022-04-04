package must

import (
	"github.com/go-redis/redis/v8"
	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
)

func ConnectRedis(cfgRedis *configs.Redis) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfgRedis.Addr,
		DB:   cfgRedis.DB,
	})

	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		ll.Fatal("Error pinging redis", l.Error(err))
	}
	return redisClient
}
