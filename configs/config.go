package configs

import (
	"os"
	"strconv"
	"time"
)

// Config service config
type Config struct {
	// Provider configs
	EncryptSalt             string
	JwtSigningKey           string
	JwtAccessTokenDuration  time.Duration
	JwtRefreshTokenDuration time.Duration
	// Database configs
	DatabaseURI   string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	// Server configs
	Host string
	Port int
}

// NewConfig constructor
func NewConfig() *Config {
	// Get provider configs
	encryptSalt := os.Getenv("ENCRYPT_SALT")
	if encryptSalt == "" {
		encryptSalt = defaultEncryptSalt
	}
	jwtSigningKey := os.Getenv("JWT_KEY")
	if jwtSigningKey == "" {
		jwtSigningKey = defaultJwtSigningKey
	}
	jwtAccessTokenDuration, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_DURATION"))
	if err != nil {
		jwtAccessTokenDuration = defaultJwtAccessTokenDuration
	}
	jwtRefreshTokenDuration, err := time.ParseDuration(os.Getenv("JWT_REFRESH_TOKEN_DURATION"))
	if err != nil {
		jwtAccessTokenDuration = defaultJwtAccessTokenDuration
	}
	// Get database configs
	databaseURI := os.Getenv("DB_URL")
	if databaseURI == "" {
		databaseURI = defaultDatabaseURI
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = defaultRedisAddr
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}
	// Get server configs
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = defaultPort
	}
	return &Config{
		EncryptSalt:             encryptSalt,
		JwtSigningKey:           jwtSigningKey,
		JwtAccessTokenDuration:  jwtAccessTokenDuration,
		JwtRefreshTokenDuration: jwtRefreshTokenDuration,
		DatabaseURI:             databaseURI,
		RedisAddr:               redisAddr,
		RedisPassword:           redisPassword,
		RedisDB:                 redisDB,
		Host:                    host,
		Port:                    port,
	}
}
