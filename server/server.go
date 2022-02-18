package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/db"
)

// APIServer is the server for http
type APIServer struct {
	router     *gin.Engine
	httpServer *http.Server
}

// NewAPIServer is the constructor for APIServer
func NewAPIServer(dbs *db.DB) *APIServer {

	// Initialize the routes
	routes := NewAPIRoutes(dbs)
	routes.SetRoutes()

	// Set the server config
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: routes.GetEngine(),
	}

	return &APIServer{
		router:     routes.GetEngine(),
		httpServer: httpServer,
	}
}

// Run serve the http server
func (server *APIServer) Run() error {
	if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown shutdows the http server
func (server *APIServer) Shutdown(ctx context.Context) error {
	if err := server.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
