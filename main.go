package main

import (
	"togo/src"
	"togo/src/api/http"
	"togo/src/infra/service"

	"github.com/golobby/container"
)

func main() {
	container.Singleton(func() src.IContextService {
		return service.NewServiceContext()
	})
	webServer := http.NewWebServer()
	webServer.LoadRouterV1().Start()
}
