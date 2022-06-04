package router

import "net/http"

type RouterInterface interface {
	Router() http.Handler
}
