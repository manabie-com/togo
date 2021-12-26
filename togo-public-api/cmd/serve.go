package cmd

import (
	"fmt"
	"net/http"
	"os/signal"
	"strconv"

	"os"

	"togo-public-api/internal/auth"
	"togo-public-api/internal/handler"
	"togo-public-api/internal/service/togo_internal_v1"
	"togo-public-api/internal/service/togo_user_session_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/grpc_client"
	"github.com/giahuyng98/togo/core-lib/grpc_gw_server"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the togo-public-api server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	},
}

func run() error {

	togoInternalConn, err := grpc_client.NewConnection(viper.GetString("services.togo_internal_service"))

	if err != nil {
		return fmt.Errorf("error connect to togo-internal-service %v", err)
	}

	togoUserSessionConn, err := grpc_client.NewConnection(viper.GetString("services.togo_user_session_service"))

	if err != nil {
		return fmt.Errorf("error connect to togo-user-session-service %v", err)
	}

	h := handler.Handler{
		TogoInternalService:    togo_internal_v1.NewTogoInternalServiceClient(togoInternalConn),
		TogoUserSessionService: togo_user_session_v1.NewTogoUserSessionServiceClient(togoUserSessionConn),
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := grpc_gw_server.New(
		"togo-public-api",
		port,
		v1.RegisterTogoPublicAPIServiceHandlerFromEndpoint,
		grpc.ChainUnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.NewAuthFunc(h.TogoUserSessionService, "RegisterOrLogin"))),
		grpc.ChainStreamInterceptor(grpc_auth.StreamServerInterceptor(auth.NewAuthFunc(h.TogoUserSessionService, "RegisterOrLogin"))),
	)
	v1.RegisterTogoPublicAPIServiceServer(server, h)

	go runPrometheus()

	fmt.Println("start serving")
	if err := server.Start(); err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("catch signal")

	server.Stop()
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
