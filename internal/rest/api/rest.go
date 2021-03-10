package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/rest/middleware"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Rest is a rest access server
type Rest struct {
	Version     string
	httpServer  *http.Server
	TodoCtrl    *TodoCtrl
	HttpPort    int
	lock        sync.RWMutex
}

func (s *Rest) makeHTTPServer(port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}
func (s *Rest) routes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.RealIP)
	router.Methods(http.MethodOptions).HandlerFunc(HandleCORSHeader)
	router = s.TodoCtrl.registerRouter(router)
	return router
}

func (s *Rest) Shutdown() {
	log.Print("[WARN] shutdown rest server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.lock.Lock()
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			log.Printf("[DEBUG] http shutdown error, %s", err)
		}
		log.Print("[DEBUG] shutdown http server completed")
	}

	s.lock.Unlock()
}

func (s *Rest) Run(port int) {
	s.httpServer = s.makeHTTPServer(port, s.routes())
	s.httpServer.ErrorLog = &log.Logger{}
	log.Fatal(s.httpServer.ListenAndServe())
}

func HandleCORSHeader(writer http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	if strings.Index(origin, "localhost") > 0 {
		writer.Header().Set("Access-Control-Allow-Origin", origin)
		writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS,DELETE")
		writer.Header().Set("Access-Control-Max-Age", "86400")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, If-None-Match, Authorization, Accept, QToken, Accept-Encoding")
	}
	if strings.ToUpper(r.Method) == "OPTIONS" {
		writer.WriteHeader(204) // send the headers with a 204 response code.
		return
	}
	return
}
