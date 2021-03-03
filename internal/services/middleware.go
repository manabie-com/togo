package services

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		// Do stuff here
		log.Println(req.Method, req.URL.Path)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(resp, req)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(resp, req)
	})
}

func (s *ToDoService) authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(resp, req)
	})
}
