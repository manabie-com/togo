package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"togo/config"
	"togo/database/datastore"
	"togo/domain/registry"
)

const DEFAULT_CONFIG_PATH = "config_dev.yaml"

func main() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = DEFAULT_CONFIG_PATH
	}

	err := config.Setup(fmt.Sprintf("config/%s", DEFAULT_CONFIG_PATH))
	if err != nil {
		panic(fmt.Errorf("error getting config %v", err))
	}

	err = datastore.SetupDB()
	if err != nil {
		panic(fmt.Errorf("error connecting to db %v", err))
	}

	listener, errChan := runHTTPServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func(errChan chan error) {
		errCh := errChan
		select {
		case <-sigCh:
			_ = listener.Close()
		case _ = <-errCh:
			_ = listener.Close()
		}
		cancel()
	}(errChan)
	<-ctx.Done()
}

func runHTTPServer() (listener net.Listener, ch chan error) {
	router, err := registry.InitHTTPServer(context.Background())
	if err != nil {
		panic(err)
	}

	conf := config.GetConfig()

	addr := fmt.Sprintf(":%s", conf.Server.ServerPort)
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	ch = make(chan error)
	go func() {
		ch <- http.Serve(listener, router)
	}()

	return listener, ch
}
