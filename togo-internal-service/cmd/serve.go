package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"togo-internal-service/internal/handler"
	"togo-internal-service/internal/storage"
	"togo-internal-service/internal/storage/sqlite"
	v1 "togo-internal-service/pkg/api/v1"
	"github.com/giahuyng98/togo/core-lib/grpc_server"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the gRPC server",
	Long:  "serve the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

	},
}

func run() error {

	// setup sqlite storage
	sqlDB, err := sqlite.NewSqliteDB(
		viper.GetString("sqlite.dbname"),
		storage.StorageConfig{
			MaxTaskCreatedPerDay: viper.GetInt("setting.max_task_created_per_day"),
			SubstrContentLength:  viper.GetInt("setting.substr_content_length"),
			SonyflakeStartTime:   viper.GetTime("setting.sonyflake_start_time"),
		})

	if err != nil {
		return err
	}

	h := handler.Handler{
		Storage: sqlDB,
		Config: &handler.Config{
			MaxListTaskPageSize: viper.GetInt("setting.max_list_task_page_size"),
		},
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s := grpc_server.New("togo-internal-service", port)
	v1.RegisterTogoInternalServiceServer(s, h)

	go runPrometheus()

	fmt.Println("start serving")
	if err = s.Start(); err != nil {
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
