package http_server

import (
	"github.com/labstack/echo/v4"

	appMiddleware "github.com/manabie-com/togo/app/web/http_server/middleware"

	gRoute "github.com/manabie-com/togo/app/common/gstuff/route"
)

func (app *apiServer) initRoute(e *echo.Echo) {
	app.authenRoute(e)
	app.publicRoute(e)
}

// authenRoute support api need basic authen
func (app *apiServer) authenRoute(e *echo.Echo) {
	api := gRoute.APIRoute(e)
	{
		_ = api
	}
}

// publicRoute support api call from public, need middle authentication
func (app *apiServer) publicRoute(e *echo.Echo) {

	publicApi := gRoute.PublicAPIRoute(e).Group("", appMiddleware.CallFromPublicRoute())
	{
		// user gr
		publicApiUserGr := publicApi.Group("/user")
		{
			publicApiUserGr.POST("/create", app.userSrv.Create)
		}

		// publicAuthenTokenApi required token login of enforcer
		publicAuthenTokenApi := publicApi.Group("", appMiddleware.NewStaff().ValidateToken)
		{
			_ = publicAuthenTokenApi
		}
	}
}
