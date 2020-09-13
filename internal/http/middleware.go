package http

import "net/http"

type Middleware func(handler http.Handler) http.HandlerFunc

func Apply(h http.Handler, mws ...Middleware) http.Handler {
  for i := len(mws)-1; i >= 0; i-- {
    h = mws[i](h)
  }
  return h
}

func ApplyFunc(h http.HandlerFunc, mws ...Middleware) http.Handler {
  return Apply(h, mws...)
}
