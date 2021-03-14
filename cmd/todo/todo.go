package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manabie-com/togo/internal/pkg/db/postgres"
	mm "github.com/manabie-com/togo/internal/pkg/middleware"
	"github.com/manabie-com/togo/internal/todo/handler"
	pgr "github.com/manabie-com/togo/internal/todo/repository/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Global log
	log.SetFormatter(&log.JSONFormatter{})
	// Http log
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{
		DisableTimestamp: true,
	}
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mm.NewStructuredLogger(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.Heartbeat("/ping"))

	sqlxConn := postgres.NewSQLXConn()
	baseRepo := pgr.PGRepository{DBConn: sqlxConn}
	appHandler := handler.NewTodoHandler(handler.TodoRepositoryList{
		UserRepo: &pgr.PGUserRepository{PGRepository: baseRepo},
		TaskRepo: &pgr.PGTaskRepository{PGRepository: baseRepo},
	})
	router.Mount("/", appHandler)

	addr := ":5050"
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithFields(log.Fields{
				"err": err,
			}).Errorln("Error starting server")
		}
	}()

	defer func() {
		gracefulShutdown(server)
		sqlxConn.Close()
	}()

	log.Infof("Start server at %s", addr)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Infoln("os.Interrupt - Shutting Down")
}

func gracefulShutdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatalln("Could not shutdown server correctly")
	} else {
		log.Infoln("Server gracefully stopped")
	}
}
