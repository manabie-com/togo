package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/jwt"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, userid, ok := jwt.CheckToken(w, r)
        if !ok {
            http.Error(w, "you need to login first to perform this action", http.StatusUnauthorized)
            return
        }
        context.Set(r, "userid", userid)
        next.ServeHTTP(w, r)
    })
}
func MiddlewareID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, err := strconv.Atoi(params["id"])
        if err != nil {
            http.Error(w, "id url need to be a number", http.StatusBadRequest)
            return
        }
        context.Set(r, "id", id)
        next.ServeHTTP(w, r)
    })
}