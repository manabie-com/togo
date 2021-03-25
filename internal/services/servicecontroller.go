package services

import (
	"net/http"

	configurations "github.com/manabie-com/togo/internal/configurations"
	"github.com/manabie-com/togo/internal/middleware"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

type ServiceController struct {
	Config configurations.Config
	Store  postgres.Store
}

func (sc *ServiceController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/login":
		middleware.NewPublicHandler(sc.loginHandler).DoFilter(resp, req)
		return
	case "/tasks":
		switch req.Method {
		case http.MethodGet:
			middleware.NewAuthHandler(sc.Config.JWTKey, sc.taskListHandler).DoFilter(resp, req)
		case http.MethodPost:
			middleware.NewAuthHandler(sc.Config.JWTKey, sc.updateTasksHandler).DoFilter(resp, req)
		}
		return
	}
}
