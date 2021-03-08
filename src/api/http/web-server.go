package http

import (
	"net/http"
	"os"
	"togo/src"
	v1 "togo/src/api/http/v1"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	frameWork *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (ws *WebServer) LoadRouterV1() src.IWebServer {
	ws.frameWork.Validator = &CustomValidator{
		validator: validator.New(),
	}
	routerV1 := v1.NewRouterV1(ws.frameWork)
	routerV1.LoadAPI()
	return ws
}

func (ws *WebServer) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := ws.frameWork.Start(":" + port)
	if err != nil {
		ws.frameWork.Logger.Fatal(err)
	}
}

func NewWebServer() src.IWebServerSetup {
	return &WebServer{
		echo.New(),
	}
}
