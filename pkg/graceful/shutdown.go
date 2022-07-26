package graceful

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DefaultShutdownTimeout = 5 * time.Second

type (
	StartFunc    func() error
	ShutdownFunc func(context.Context) error
)

func Graceful(start StartFunc, shutdown ShutdownFunc) error {
	var (
		stopChan = make(chan os.Signal)
		errChan  = make(chan error)
	)

	go wait(stopChan, errChan, shutdown)

	if err := start(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}

func wait(stopChan chan os.Signal, errChan chan error, shutdown ShutdownFunc) {
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), DefaultShutdownTimeout)
	defer cancel()

	if err := shutdown(ctx); err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
