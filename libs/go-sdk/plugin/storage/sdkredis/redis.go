package sdkredis

import (
	"context"
	"flag"

	"github.com/go-redis/redis/v8"
	"github.com/phathdt/libs/go-sdk/logger"
)

var (
	defaultRedisName      = "DefaultRedis"
	defaultRedisMaxActive = 0 // 0 is unlimited max active connection
	defaultRedisMaxIdle   = 10
)

type RedisDBOpt struct {
	Prefix    string
	RedisUri  string
	MaxActive int
	MaxIde    int
}

type redisDB struct {
	name   string
	client *redis.Client
	logger logger.Logger
	*RedisDBOpt
}

func NewRedisDB(name, flagPrefix string) *redisDB {
	return &redisDB{
		name: name,
		RedisDBOpt: &RedisDBOpt{
			Prefix:    flagPrefix,
			MaxActive: defaultRedisMaxActive,
			MaxIde:    defaultRedisMaxIdle,
		},
	}
}

func (r *redisDB) GetPrefix() string {
	return r.Prefix
}

func (r *redisDB) isDisabled() bool {
	return r.RedisUri == ""
}

func (r *redisDB) InitFlags() {
	prefix := r.Prefix
	if r.Prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&r.RedisUri, prefix+"uri", "redis://localhost:6379", "(For go-redis) Redis connection-string. Ex: redis://localhost/0")
	flag.IntVar(&r.MaxActive, prefix+"pool-max-active", defaultRedisMaxActive, "(For go-redis) Override redis pool MaxActive")
	flag.IntVar(&r.MaxIde, prefix+"pool-max-idle", defaultRedisMaxIdle, "(For go-redis) Override redis pool MaxIdle")
}

func (r *redisDB) Configure() error {
	if r.isDisabled() {
		return nil
	}

	r.logger = logger.GetCurrent().GetLogger(r.name)
	r.logger.Info("Connecting to Redis at ", r.RedisUri, "...")

	opt, err := redis.ParseURL(r.RedisUri)

	if err != nil {
		r.logger.Error("Cannot parse Redis ", err.Error())
		return err
	}

	opt.PoolSize = r.MaxActive
	opt.MinIdleConns = r.MaxIde

	client := redis.NewClient(opt)

	// Ping to test Redis connection
	if err = client.Ping(context.Background()).Err(); err != nil {
		r.logger.Error("Cannot connect Redis. ", err.Error())
		return err
	}

	// Connect successfully, assign client to goRedisDB
	r.client = client
	return nil
}

func (r *redisDB) Name() string {
	return r.name
}

func (r *redisDB) Get() interface{} {
	return r.client
}

func (r *redisDB) Run() error {
	return r.Configure()
}

func (r *redisDB) Stop() <-chan bool {
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			r.logger.Info("cannot close ", r.name)
		}
	}

	c := make(chan bool)
	go func() { c <- true }()
	return c
}
