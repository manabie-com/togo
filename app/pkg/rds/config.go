package rds

const (
	DsnFormat = "%s:%d"
	DefaultDb = 0
)

type RedisConfig struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}
