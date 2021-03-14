package main

import (
	"database/sql"
	"github.com/banhquocdanh/togo/internal/cache"
	"github.com/banhquocdanh/togo/internal/config"
	server2 "github.com/banhquocdanh/togo/internal/server"
	"github.com/banhquocdanh/togo/internal/services"
	sqllite "github.com/banhquocdanh/togo/internal/storages/sqlite"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	var cfg = config.Config{
		JwtKey: "wqGyEBBfPK9w3Lxw",
		Redis: config.RedisConfig{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		},
		Database: config.DatabaseConfig{
			DriverName:     "sqlite3",
			DataSourceName: "./data.db",
		},
		TokenTIL: 15,
	}
	//TODO: read config from env

	db, err := sql.Open(cfg.Database.DriverName, cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	server := server2.NewToDoHttpServer(
		cfg.JwtKey,
		services.NewToDoService(
			services.WithConfig(&cfg),
			services.WithStore(&sqllite.LiteDB{DB: db}),
			services.WithCache(cache.NewRedisCache(redisClient)),
		),
	)

	if err := server.Listen(5050); err != nil {
		panic(err)
	}

}
