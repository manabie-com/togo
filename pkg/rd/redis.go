package pkg_rd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"log"
	"manabie-com/togo/global"
	"sync"
	"time"
)

var rdProvider RedisProvider

func InitializeRdV8() {
	SetRedisProvider(&RedisV8{
		Host:        global.Config.RedisConnectionHost,
		Password:    global.Config.RedisConnectionPassword,
		DB:         1,
	})

}

type RedisProvider interface {
	RdConn() *redis.Client
}

func SetRedisProvider(provider RedisProvider) {
	rdProvider = provider
}

func HasRedisProvider() bool {
	return rdProvider != nil
}

func RdConn() *redis.Client {
	return rdProvider.RdConn()
}

type RedisV8 struct {
	Host     string
	Password string
	DB       int

	once       sync.Once
	clientConn *redis.Client
}

func (i *RedisV8) RdConn() *redis.Client {
	i.once.Do(i.Connect)

	if i.clientConn == nil {
		log.Fatal("redis not connected")
	}

	return i.clientConn
}

func (i *RedisV8) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.RedisConnectionHost,
		Password: global.Config.RedisConnectionPassword,
		DB:       1,
	})

	pong, err := client.Ping(ctx).Result()
	color.Green("Ping Connection to Redis Database: %s %v", pong, err)

	if err != nil || client == nil {
		fmt.Println("Redis connect:", err)
		log.Fatal("redis connect: ", err)
	}
	i.clientConn = client
}

func (i *RedisV8) Close() {
	if i.clientConn != nil {
		_ = i.clientConn.Close()

		i.clientConn = nil
	}
}
