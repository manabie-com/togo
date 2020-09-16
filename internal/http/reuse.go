package http

import "net/http"

var ListenAndServe = http.ListenAndServe

type ResponseWriter = http.ResponseWriter
type Request = http.Request
