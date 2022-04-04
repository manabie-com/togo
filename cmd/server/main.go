package main

import (
	"context"
	"fmt"
	"net"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/internal/must"
	"github.com/vchitai/togo/internal/server"
	"github.com/vchitai/togo/internal/service"
	"github.com/vchitai/togo/internal/store"
)

var (
	ll = l.New()
)

func main() {
	// load configs
	cfg := configs.Load()

	ctx, ctxCancel := context.WithCancel(context.Background())
	var teardownTimeout time.Duration = 15
	if cfg.Environment == "D" {
		teardownTimeout = 1
	}
	go waitForShutdownSignal(teardownTimeout, ctxCancel)

	var (
		db       = must.ConnectMySQL(cfg.MySQL)
		redisCli = must.ConnectRedis(cfg.Redis)
	)
	_ = db.AutoMigrate(&models.ToDoConfig{})
	_ = db.AutoMigrate(&models.ToDo{})
	var (
		todoStore = store.NewToDo(db, redisCli)
		svc       = service.New(cfg, todoStore)
		srv       = server.NewGRPCServer().
				WithServiceServer(svc)
		gw = server.NewGatewayServer(cfg)
	)

	go func() {
		defer ctxCancel()
		ll.Info("HTTP server start listening", l.Int("HTTP address", cfg.HTTPAddress))
		if err := gw.RunGRPCGateway(svc); err != nil {
			ll.Panic("error listening to address", l.Int("address", cfg.HTTPAddress), l.Error(err))
			return
		}
	}()

	go func() {
		defer ctxCancel()
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
		if err != nil {
			ll.Panic("error listening to address", l.Int("address", cfg.GRPCAddress), l.Error(err))
			return
		}
		ll.Info("GRPC server start listening", l.Int("GRPC address", cfg.GRPCAddress))
		_ = srv.Serve(listener)
	}()
	var shutdownWg sync.WaitGroup
	shutdownWg.Add(1)
	go func() {
		defer shutdownWg.Done()
		<-ctx.Done()
		_ = gw.Shutdown(context.TODO()) // temp no timeout
	}()
	shutdownWg.Add(1)
	go func() {
		defer shutdownWg.Done()
		<-ctx.Done()
		srv.GracefulStop()
	}()
	shutdownWg.Wait() // wait for all main process tore down
}

func waitForShutdownSignal(timeout time.Duration, callback func()) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	<-osSignal
	// Wait for maximum 15s
	go func() {
		timer := time.NewTimer(timeout)
		<-timer.C
		ll.Fatal("Force shutdown due to timeout!")
	}()
	callback()
}
