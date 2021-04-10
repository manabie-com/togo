package infra

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"go.opencensus.io/plugin/ochttp"
)

const (
	DEFAULT_PORT = "3000"
)

type RestAPIHandler http.Handler

type RestService struct {
	isStarted bool
	srv       http.Server
}

func ProvideRestService(cfg *AppConfig, handler RestAPIHandler) (*RestService, func(), error) {
	port := strconv.Itoa(cfg.Port)
	if port == "" {
		port = DEFAULT_PORT
	}

	och := &ochttp.Handler{
		Handler: handler,
	}

	rest := &RestService{
		isStarted: false,
		srv: http.Server{
			Addr:    ":" + port,
			Handler: och,

			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 2 * time.Minute,
			ReadTimeout:  2 * time.Minute,
		},
	}

	logrus.Info("Init rest APIs client completed")
	return rest, func() {
		rest.Close()
	}, nil
}

func (s *RestService) MustStart() {
	s.isStarted = true
	if err := s.srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic("Error starting server! " + err.Error())
		} else {
			fmt.Println("Shutting down...")
		}
	}
}

func (s *RestService) Close() {
	if s.isStarted {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := s.srv.Shutdown(ctx); err != nil {
			panic("Server did not shut down before timeout: " + err.Error())
		} else {
			fmt.Println("Server shutdown")
		}
	}
}
