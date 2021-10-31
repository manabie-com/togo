package settings

// ENV ...
const (
	DbDriver = "DB_DRIVER"
	PgHost   = "PG_HOST"
	PgPort   = "PG_PORT"
	PgUser   = "PG_USER"
	PgPass   = "PG_PASS"
	PgDB     = "PG_DB"

	SaltRounds = "SALT_ROUNDS"

	RedisHost           = "REDIS_HOST"
	RedisPort           = "REDIS_PORT"
	RedisCacheExpiresIn = "REDIS_CACHE_EXPIRES_IN"

	ENV         = "ENV"
	ENVFilePath = "src/pkgs/settings/config.env"
	Port        = "PORT"
)
