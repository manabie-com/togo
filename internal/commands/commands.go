package commands

import (
	"database/sql"
	"fmt"
	"github.com/banhquocdanh/togo/internal/cache"
	"github.com/banhquocdanh/togo/internal/config"
	server2 "github.com/banhquocdanh/togo/internal/server"
	"github.com/banhquocdanh/togo/internal/services"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/banhquocdanh/togo/internal/storages/postgresql"
	sqllite "github.com/banhquocdanh/togo/internal/storages/sqlite"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
)

func ToDoCommand() (*cobra.Command, error) {
	var port int
	var useSqlLite bool

	var todoCommand = &cobra.Command{
		Use:   "todo",
		Short: "start todo service",
		Long:  "start todo service",
		Run: func(cmd *cobra.Command, args []string) {
			var cfg = config.Config{}
			err := config.LoadConfigFromEnv(&cfg)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Cfg: %+v\n", cfg)
			redisClient := redis.NewClient(&redis.Options{
				Addr:     cfg.Redis.Addr,
				Password: cfg.Redis.Password,
				DB:       cfg.Redis.DB,
			})
			var store storages.StoreInterface
			if useSqlLite == true {
				log.Printf("used sqlite database")
				db, err := sql.Open("sqlite3", "./data.db")
				if err != nil {
					panic(err)
				}
				store = sqllite.NewSqlLite(db)
			} else {
				log.Printf("used postgres database")
				redSyncClient := redsync.New(goredis.NewPool(redisClient))
				store = postgresql.NewPostgreSQL(
					pg.Connect(&pg.Options{
						Addr:     cfg.Database.Addr,
						User:     cfg.Database.User,
						Password: cfg.Database.Password,
						Database: cfg.Database.Database,
					}),
					postgresql.WithRedSync(redSyncClient),
				)
			}

			server := server2.NewToDoHttpServer(
				cfg.JwtKey,
				services.NewToDoService(
					services.WithConfig(&cfg),
					services.WithStore(store),
					services.WithCache(cache.NewRedisCache(redisClient)),
				),
			)
			if err := server.Listen(5050); err != nil {
				panic(err)
			}
		},
	}

	todoCommand.Flags().IntVarP(&port, "port", "P", 5050, "listen on port")
	todoCommand.Flags().BoolVarP(&useSqlLite, "sqlLite", "l", false, "use sqlLite database, default use postgres")

	return todoCommand, nil
}
