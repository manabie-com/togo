package main

import "net/http"

func DefaultMiddleWare(next http.Handler) http.Handler {
	return nil
}
