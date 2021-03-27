package servehttp

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type AppServer struct {
	router *mux.Router
}

func (a *AppServer) Init() {
	a.router = mux.NewRouter().StrictSlash(true)

	// Register common handlers
	a.router.NotFoundHandler = &NotFoundHandler{}
	a.router.MethodNotAllowedHandler = &MethodNotAllowedHandler{}
}

func (a *AppServer) RegisterHandler(method string, route string, handler IAPIHandler) {
	if len(method) == 0 {
		method = http.MethodGet
	}
	h := logHandler(handler.ServeHTTP)

	a.router.HandleFunc(route, h.ServeHTTP).Methods(method)
}

func (a *AppServer) GetRouter() *mux.Router {
	return a.router
}

// Log request handler
func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("[%s] Host: %s, Uri: %s", r.Method, r.Host, r.RequestURI))
		rec := httptest.NewRecorder()
		tStart := time.Now()
		fn(rec, r)
		log.Println(fmt.Sprintf("[%d] Ext: %v, Response: %s", rec.Code, time.Now().Sub(tStart), rec.Body))

		// copies the recorded response to the response writer
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)
	}
}

// Common handlers
type NotFoundHandler struct {
}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusNotFound, map[string]string{
		"error": "Not found.",
	})
}

type MethodNotAllowedHandler struct {
}

func (MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusMethodNotAllowed, map[string]string{
		"error": "Method not allowed.",
	})
}
