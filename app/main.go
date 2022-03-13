package main

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/triet-truong/todo/config"
	"github.com/triet-truong/todo/servers"
	"github.com/triet-truong/todo/todo/repository"
)

func init() {
	// Load env vars
	config.Load()
}

func main() {
	// Setup repositories
	repo := repository.NewTodoMysqlRepository(config.DatabaseDSN())
	cacheStore := repository.NewTodoRedisRepository(redis.Options{
		Addr:     config.CacheConnectioURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//Run HTTP server
	server := servers.NewServer(repo, cacheStore)
	server.Run()
}
