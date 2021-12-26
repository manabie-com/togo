package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"togo-user-session-service/internal/handler"
	"togo-user-session-service/internal/storage/combined"
	"togo-user-session-service/internal/storage/redis"
	"togo-user-session-service/internal/storage/sqlite"
	v1 "togo-user-session-service/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/grpc_server"
	"github.com/giahuyng98/togo/core-lib/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start serve user-session-service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	},
}

func run() error {

  logger.Info("creating redis database")
	redisDB, err := redis.NewRedisStorage(
		&redis.Config{
			Addr:      viper.GetString("redis.address"),
			Password:  viper.GetString("redis.password"),
			DB:        viper.GetInt("redis.db"),
			TokenTTL:  viper.GetDuration("setting.token_ttl"),
			SecretKey: viper.GetString("setting.secret_key"),
		},
	)
  logger.Info("done creating redis database")

	if err != nil {
		return err
	}

  logger.Info("creating sqlite database")
	sqliteDB, err := sqlite.NewSqliteDB(
		&sqlite.Config{
			Name:               viper.GetString("sqlite.dbname"),
			SonyflakeStartTime: viper.GetTime("setting.sonyflake_start_time"),
		},
	)
  logger.Info("done creating sqlite database")

	if err != nil {
		return err
	}

	h := handler.Handler{
		DB: &combined.CombinedDB{
			SessionDB: redisDB,
			UserDB:    sqliteDB,
		},
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s := grpc_server.New("togo-user-session-service", port)
	v1.RegisterTogoUserSessionServiceServer(s, h)

	go runPrometheus()

	fmt.Println("start serving")
	if err := s.Start(); err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("catch signal")

	s.Stop()

	return nil
}

func runPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	metricPort, _ := strconv.Atoi(os.Getenv("METRIC_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", metricPort), nil); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
