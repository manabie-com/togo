package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewConfig(httpPort int) Config {
	return Config{
		HTTP: ServerListen{
			Host: "0.0.0.0",
			Port: httpPort,
		},
	}
}

type Config struct {
	HTTP ServerListen
}

type ServerListen struct {
	Host string
	Port int
}

func (l ServerListen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

type Server struct {
	cfg     Config
	handler http.Handler
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Register(handler http.Handler) {
	s.handler = handler
}

func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	httpServer := http.Server{
		Addr:    s.cfg.HTTP.String(),
		Handler: s.handler,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			errch <- err
		}
	}()

	for {
		select {
		case <-stop:
			ctx, cancelFn := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancelFn()
			if err := httpServer.Shutdown(ctx); err != nil {
				// log.Errorf("failed to stop server: %w", err)
			}
			return nil
		case err := <-errch:
			return err
		}
	}
}
