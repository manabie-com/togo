package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
    beego.GlobalControllerRouter["togo/controllers:HealthCheckController"] = append(beego.GlobalControllerRouter["togo/controllers:HealthCheckController"],
        beego.ControllerComments{
            Method:           "Get",
            Router:           "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams:     param.Make(),
            Filters:          nil,
            Params:           nil})

    beego.GlobalControllerRouter["togo/controllers:TaskController"] = append(beego.GlobalControllerRouter["togo/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
