package controllers

import beego "github.com/beego/beego/v2/server/web"

type HealthCheckController struct {
	beego.Controller
}

// @Title Get
// @router / [get]
func (m *HealthCheckController) Get() {
	m.Data["json"] = beego.BConfig.AppName
	defer m.Ctx.Request.Body.Close()
	_ = m.ServeJSON()
}