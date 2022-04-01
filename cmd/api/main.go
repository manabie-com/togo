package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", healthCheck)
	e.Server.Addr = fmt.Sprintf(":8080")
	e.Server.ReadTimeout = time.Duration(10) * time.Second
	e.Server.WriteTimeout = time.Duration(5) * time.Second

	e.StartServer(e.Server)
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
