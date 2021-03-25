package middleware

import "net/http"

type UserAuthKey int8

type HandlerFunc func(http.ResponseWriter, *http.Request)
