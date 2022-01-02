package adapter

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Redises ..
type Redises map[string]*Redis

// Get Redis
func (adapters Redises) Get(name string) (result *Redis) {
	if adapter, ok := adapters[name]; ok {
		result = adapter
	} else {
		panic("Không tìm thấy config Redis " + name)
	}
	return
}

// Redis ..
type Redis struct {
	Name           string        `mapstructure:"name"`
	ConnectionType string        `mapstructure:"connection-type"`
	Address        string        `mapstructure:"address"`
	Password       string        `mapstructure:"password"`
	DBNum          int           `mapstructure:"dbnum"`
	Retry          int           `mapstructure:"retry"`
	Timeout        time.Duration `mapstructure:"timeout"`
	PoolLimit      int           `mapstructure:"pool-limit"`
	Client         *redis.Client
}

var (
	onceRedis      map[string]*sync.Once
	onceRedisMutex = sync.RWMutex{}
)

func init() {
	onceRedis = make(map[string]*sync.Once)
}

// Init ..
func (config *Redis) Init() {
	onceRedisMutex.Lock()

	if onceRedis[config.Name] == nil {
		onceRedis[config.Name] = &sync.Once{}
	}
	var connectError error
	onceRedis[config.Name].Do(func() {
		log.Printf("[%s][%s] Redis [connecting]\n", config.Name, config.Address)
		config.Client = redis.NewClient(&redis.Options{
			Network: config.ConnectionType,
			Addr:    config.Address,
			OnConnect: func(*redis.Conn) error {
				log.Printf("[%s][%s] Redis [connected]\n", config.Name, config.Address)
				return nil
			},
			Password:    config.Password,
			DB:          config.DBNum,
			MaxRetries:  config.Retry,
			DialTimeout: config.Timeout * time.Second,
			PoolSize:    config.PoolLimit,
		})
		_, err := config.Client.Ping().Result()
		if err != nil {
			connectError = err
		}
	})

	onceRedisMutex.Unlock()

	if connectError != nil {
		log.Printf("Could not connect to redis %v\n", connectError)
		time.Sleep(1 * time.Second)
		onceRedis[config.Name] = &sync.Once{}
		config.Init()
		return
	}
}

// GetClient func
func (config Redis) GetClient() (client *redis.Client) {
	if config.Client != nil {
		client = config.Client
	} else {
		panic(fmt.Errorf("[%s] Chưa init Redis", config.Name))
	}
	return
}
