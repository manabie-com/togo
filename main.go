package main

import (
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/data/repository"
	"github.com/manabie-com/togo/handlers"
	"github.com/manabie-com/togo/infra"
	"github.com/manabie-com/togo/pkg/core"
	"github.com/manabie-com/togo/pkg/core/servehttp"
	"github.com/manabie-com/togo/usecases"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Init App server
	appServer := &servehttp.AppServer{}
	appServer.Init()

	db, err := infra.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	appHandlers, err := getListHandler(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, handler := range appHandlers {
		appServer.RegisterHandler(handler.Method, handler.Route, handler.Handler)
	}

	localPort := os.Getenv("APP_PORT")
	srv := &http.Server{
		Handler:      appServer.GetRouter(),
		Addr:         fmt.Sprintf(":%v", localPort),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println(fmt.Sprintf("Starting API server with port :%v", localPort))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	servehttp.WaitForShutdown(srv)
}

// Defines list handler to serve requests
func getListHandler(db *sql.DB) ([]servehttp.AppHandler, error) {

	userRepo := repository.NewUserRepositoryImpl(db)
	taskRepo := repository.NewTaskRepositoryImpl(db)
	auth, err := core.NewAppAuthenticator()

	if err != nil {
		return nil, err
	}

	return []servehttp.AppHandler{
		{
			Route:   "/login",
			Method:  http.MethodPost,
			Handler: &handlers.LoginHandler{
				Uc: usecases.NewLoginUseCase(userRepo, auth),
			},
		},

		// ======================================
		{
			Route:  "/tasks",
			Method: http.MethodGet,
			Handler: &handlers.GetTasksHandler{
				Uc: usecases.NewGetTasksUseCase(taskRepo),
				Auth: auth,
			},
		},
		// Create task
		{
			Route:  "/tasks",
			Method: http.MethodPost,
			Handler: &handlers.CreateTaskHandler{
				Uc: usecases.NewCreateTaskUseCase(taskRepo, userRepo),
				Auth: auth,
			},
		},
	}, nil
}
