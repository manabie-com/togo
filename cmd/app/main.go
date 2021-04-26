package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	nethttp "net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/storages/psql"
	"github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	conf struct {
		Transport  http.APIConf
		AuthUC     usecase.AuthUCConf `envconfig:"AUTH"`
		Storage    psql.Config
		Port       int
		MetricPort int
	}
)

func init() {
	err := envconfig.Process("app", &conf)
	if err != nil {
		panic(err)
	}
}

func setupPrometheus() {
	nethttp.Handle("/metrics", promhttp.Handler())
	nethttp.ListenAndServe(fmt.Sprintf(":%d", conf.MetricPort), nil)
}

func main() {
	go setupPrometheus()
	e := echo.New()
	storage, err := psql.NewStorage(conf.Storage)
	if err != nil {
		panic(err)
	}

	taskUc := usecase.NewTaskUseCase(storage, storage)
	authUc, err := usecase.NewAuthUseCase(conf.AuthUC, storage)
	if err != nil {
		panic(err)
	}
	http.BindAPI(conf.Transport, e, taskUc, authUc)

	go func() {
		e.Start(fmt.Sprintf(":%d", conf.Port))
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	e.Shutdown(context.Background())
}
