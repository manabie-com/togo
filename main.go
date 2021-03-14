package main

import (
	"github.com/banhquocdanh/togo/internal/cache"
	"github.com/banhquocdanh/togo/internal/config"
	server2 "github.com/banhquocdanh/togo/internal/server"
	"github.com/banhquocdanh/togo/internal/services"
	"github.com/banhquocdanh/togo/internal/storages/postgresql"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
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
			Addr:     "127.0.0.1:5432",
			User:     "bandan",
			Password: "bandan",
			Database: "todo",
		},
		TokenTIL: 15,
	}
	//TODO: read config from env

	//db, err := sql.Open("sqlite3", "./data.db")
	//if err != nil {
	//	log.Fatal("error opening db", err)
	//}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	pgStore := postgresql.NewPostgreSQL(pg.Connect(&pg.Options{
		Addr:     cfg.Database.Addr,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.Database,
	}))

	server := server2.NewToDoHttpServer(
		cfg.JwtKey,
		services.NewToDoService(
			services.WithConfig(&cfg),
			services.WithStore(pgStore),
			services.WithCache(cache.NewRedisCache(redisClient)),
		),
	)

	if err := server.Listen(5050); err != nil {
		panic(err)
	}

}
