package main

import (
	"fmt"

	"github.com/TrinhTrungDung/togo/config"
	"github.com/TrinhTrungDung/togo/internal/api/auth"
	"github.com/TrinhTrungDung/togo/internal/api/plan"
	"github.com/TrinhTrungDung/togo/internal/api/subscription"
	"github.com/TrinhTrungDung/togo/pkg/crypter"
	"github.com/TrinhTrungDung/togo/pkg/db"
	"github.com/TrinhTrungDung/togo/pkg/jwt"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSslMode), cfg.DbLog)
	if err != nil {
		panic(err)
	}

	// Initialize HTTP server
	e := server.New(&server.Config{
		Port:         cfg.ServerPort,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
		Debug:        cfg.ServerDebug,
	})

	// Initialize necessary services
	crypterSvc := crypter.New()
	jwtSvc := jwt.New(cfg.JwtAlgorithm, cfg.JwtSecret, cfg.JwtDuration)
	authSvc := auth.New(db, crypterSvc, jwtSvc)
	planSvc := plan.New(db)
	subscriptionSvc := subscription.New(db)

	// Initialize root API
	rootRouter := e.Group("/api")
	auth.NewHTTP(authSvc, rootRouter.Group("/auth"))
	plan.NewHTTP(planSvc, rootRouter.Group("/plans"))
	subscription.NewHTTP(subscriptionSvc, authSvc, rootRouter.Group("/subscriptions", jwtSvc.MWFunc()))

	// Start the HTTP server
	server.Start(e)
}
