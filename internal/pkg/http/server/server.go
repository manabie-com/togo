package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dinhquockhanh/togo/internal/pkg/config"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
)

func Start(conf *config.Server, router http.Handler) {
	addr := fmt.Sprintf("%s:%s", conf.Addr, conf.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       *conf.ReadTimeout,
		ReadHeaderTimeout: *conf.ReadHeaderTimeout,
		WriteTimeout:      *conf.WriteTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("listen and serve: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)

	// kill -2 is syscall.SIGINT
	// kill (no param) default send syscall.SIGTERM

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof(
		"Shutting down server with %.fs timeout.",
		conf.ShutdownTimeOut.Seconds(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), *conf.ShutdownTimeOut)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("Server forced to shutdown: %s", err)
	}

	log.Info("Server exited.")
}
