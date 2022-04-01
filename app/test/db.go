package test

import (
	"ansidev.xyz/pkg/rds"
	"os"
	"strconv"

	"ansidev.xyz/pkg/db"
)

func GetTestDbConfig() (db.SqlDbConfig, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	if err != nil {
		return db.SqlDbConfig{}, err
	}

	return db.SqlDbConfig{
		DbDriver:   dbDriver,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUsername: dbUsername,
		DbPassword: dbPassword,
	}, nil
}

func GetTestRedisDbConfig() (rds.RedisConfig, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if err != nil {
		return rds.RedisConfig{}, err
	}

	return rds.RedisConfig{
		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,
	}, nil
}
