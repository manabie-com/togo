package config

import (
	"github.com/ansidev/togo/constant"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) SetupTest() {
	os.Unsetenv("APP_ENV")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("HOST")
	os.Unsetenv("PORT")
	os.Unsetenv("DB_DRIVER")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("REDIS_HOST")
	os.Unsetenv("REDIS_PORT")
	os.Unsetenv("REDIS_PASSWORD")
	os.Unsetenv("TOKEN_TTL")
}

func (s *ConfigTestSuite) TestLoadConfigProd() {
	os.Setenv("APP_ENV", constant.DefaultProdEnv)
	os.Setenv("LOG_LEVEL", "error")
	os.Setenv("HOST", "https://github.com")
	os.Setenv("PORT", "80")
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "demo")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("REDIS_HOST", "127.0.0.1")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("REDIS_PASSWORD", "prod_p4ssw0rd")
	os.Setenv("TOKEN_TTL", "600")

	var config Config
	LoadConfig("..", "app.env.example", &config)

	assert.Equal(s.T(), os.Getenv("APP_ENV"), constant.DefaultProdEnv, "APP_ENV should be prod")
	assert.Equal(s.T(), "error", config.LogLevel, "LogLevel should be error")

	assert.Equal(s.T(), "https://github.com", config.Host, "Host should be https://github.com")
	assert.Equal(s.T(), 80, config.Port, "Port should be 80")

	assert.Equal(s.T(), "postgres", config.DbDriver, "DbHost should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.DbHost, "DbHost should be 127.0.0.1")
	assert.Equal(s.T(), 5432, config.DbPort, "DbPort should be 5432")
	assert.Equal(s.T(), "demo", config.DbName, "DbName should be demo")
	assert.Equal(s.T(), "postgres", config.DbUsername, "DbUsername should be postgres")
	assert.Equal(s.T(), "postgres", config.DbPassword, "DbPassword should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.RedisHost, "RedisHost should be 127.0.0.1")
	assert.Equal(s.T(), 6379, config.RedisPort, "RedisPort should be 6379")
	assert.Equal(s.T(), "prod_p4ssw0rd", config.RedisPassword, "RedisPassword should be prod_p4ssw0rd")
	assert.Equal(s.T(), 600, config.TokenTTL, "TokenTTL should be 600")
}

func (s *ConfigTestSuite) TestLoadConfigFromEnvFile() {
	os.Setenv("APP_ENV", "local")

	var config Config
	LoadConfig("..", "app.env.example", &config)

	assert.Equal(s.T(), os.Getenv("APP_ENV"), "local", "APP_ENV should be local")
	assert.Equal(s.T(), "debug", config.LogLevel, "LogLevel should be debug")

	assert.Equal(s.T(), "localhost", config.Host, "Host should be localhost")
	assert.Equal(s.T(), 8080, config.Port, "Port should be 8080")

	assert.Equal(s.T(), "postgres", config.DbDriver, "DbHost should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.DbHost, "DbHost should be 127.0.0.1")
	assert.Equal(s.T(), 5432, config.DbPort, "DbPort should be 5432")
	assert.Equal(s.T(), "todo", config.DbName, "DbName should be todo")
	assert.Equal(s.T(), "postgres", config.DbUsername, "DbUsername should be postgres")
	assert.Equal(s.T(), "postgres", config.DbPassword, "DbPassword should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.RedisHost, "RedisHost should be 127.0.0.1")
	assert.Equal(s.T(), 6379, config.RedisPort, "RedisPort should be 6379")
	assert.Equal(s.T(), "p4ssw0rd", config.RedisPassword, "RedisPassword should be p4ssw0rd")
	assert.Equal(s.T(), 86400, config.TokenTTL, "TokenTTL should be 86400")
}
