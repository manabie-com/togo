package middleware

import "net/http"

type CORS struct{}

func (c CORS) All(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")

    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusOK)
      return
    }

    next.ServeHTTP(w, r)
  }
}
