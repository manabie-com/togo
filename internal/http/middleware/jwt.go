package middleware

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "net/http"
  "strings"
)

type JWT struct {
  JwtKey              string
  UserRepo            core.UserRepo
  UnauthorizedRespond func(w http.ResponseWriter, message string)
}

func (mw *JWT) SetUser(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("Authorization")
    if err != nil {
      log.Printf("[http::JWT::SetUser - error getting cookie: %v]\n", err)
      w.WriteHeader(http.StatusUnauthorized)
      mw.UnauthorizedRespond(w, "token not found")
      return
    }
    token := strings.TrimSpace(cookie.Value)
    //token := strings.TrimSpace(r.Header.Get("Authorization"))
    claims := make(jwt.MapClaims)
    t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
      return []byte(mw.JwtKey), nil
    })
    if err != nil {
      log.Printf("[http::JWT::SetUser - parse token error: %v]\n", err)
      w.WriteHeader(http.StatusUnauthorized)
      mw.UnauthorizedRespond(w, "token invalid")
      return
    }
    if !t.Valid {
      log.Printf("[http::JWT::SetUser - invalid token: %v]\n", token)
      w.WriteHeader(http.StatusUnauthorized)
      mw.UnauthorizedRespond(w, "token invalid")
      return
    }

    id, ok := claims["user_id"].(string)
    if !ok {
      log.Printf("[http::JWT::SetUser - user_id not a string: %v]\n", id)
      w.WriteHeader(http.StatusUnauthorized)
      mw.UnauthorizedRespond(w, "token invalid")
      return
    }

    user, err := mw.UserRepo.ById(r.Context(), id)
    if err != nil {
      return
    }
    r = r.WithContext(context.WithUser(r.Context(), user))
    next.ServeHTTP(w, r)
  }
}

func (mw *JWT) RequireUser(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    user := context.User(r.Context())
    if user == nil {
      w.WriteHeader(http.StatusUnauthorized)
      mw.UnauthorizedRespond(w, "unauthorized access")
      return
    }
    next.ServeHTTP(w, r)
  }
}
