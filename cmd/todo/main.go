package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
	"github.com/oklog/run"

	httpDelivery "github.com/valonekowd/togo/adapter/delivery/http"
	"github.com/valonekowd/togo/adapter/endpoint"
	"github.com/valonekowd/togo/adapter/formatter"
	"github.com/valonekowd/togo/adapter/repository"
	"github.com/valonekowd/togo/config"
	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/infrastructure/datastore/sql"
	"github.com/valonekowd/togo/infrastructure/validator/playground"
	"github.com/valonekowd/togo/usecase"
)

func main() {
	// setup logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// load config
	cfg, err := config.Create()
	if err != nil {
		logger.Log("method", "config.Create", "err=", err)
	}

	// setup primary database
	db, err := sql.Connect(
		sql.WithDriverName(cfg.Datastore.Primary.DriverName),
		sql.WithHost(cfg.Datastore.Primary.Host),
		sql.WithPort(cfg.Datastore.Primary.Port),
		sql.WithUsername(cfg.Datastore.Primary.Username),
		sql.WithPassword(cfg.Datastore.Primary.Password),
		sql.WithDBName(cfg.Datastore.Primary.DBName),
	)
	if err != nil {
		logger.Log("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "sql.Connect", "err", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Log("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "db.Close", "err", err)
		} else {
			logger.Log("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "db.Close", "msg", "success")
		}
	}()

	// setup authentication
	jwtCfg, err := auth.NewJWTConfig(
		auth.WithKeyFunc(cfg.Auth.JWT.Secret),
		auth.WithIssuer(cfg.Auth.JWT.Issuer),
		auth.WithSigningMethod(cfg.Auth.JWT.Algorithm),
	)
	if err != nil {
		logger.Log("auth", "jwt", "method", "auth.NewJWTConfig", "err", err)
		os.Exit(1)
	}

	authCfg := auth.Config{
		JWT: jwtCfg,
		// OAuth...
	}

	// setup service
	var (
		validator   = playground.NewValidator(cfg.Validation.Playground.TagName, logger)
		repository  = repository.New(repository.WithDB(db))
		formatter   = formatter.New(authCfg)
		usecase     = usecase.New(repository, formatter, logger)
		endpoint    = endpoint.MakeServerEndpoint(usecase, authCfg, logger)
		httpHandler = httpDelivery.NewHTTPHandler(endpoint, validator, logger)
	)

	if cfg.IsProd() {
		httpHandler = http.TimeoutHandler(httpHandler, 3*time.Second, "server timeout")
	}

	var g run.Group
	// setup shutdown hook
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(error) {
			cancel()
		})
	}
	// setup server
	{
		var ln net.Listener
		g.Add(func() error {
			ln, err := net.Listen("tcp", cfg.HTTPAddr())
			if err != nil {
				return err
			}

			logger.Log("transport", "HTTP", "addr", ln.Addr().String(), "msg", "listening")
			return http.Serve(ln, httpHandler)
		}, func(error) {
			logger.Log("transport", "HTTP", "method", "net.Listen", "err", err)
			ln.Close()
		})
	}

	// start all
	logger.Log("exit", g.Run())
}
