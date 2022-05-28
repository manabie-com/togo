package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"togo/helper"
	"togo/models"
)

type TaskController struct {
	beego.Controller
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @router / [post]
func (t *TaskController) Post() {
	if accepted, code, err := helper.BeforeTask(t.Ctx); err != nil {
		t.Ctx.Output.SetStatus(code)
		t.Data["json"] = map[string]interface{}{"err": err.Error()}
	} else if accepted {
		var task models.Task
		if err = json.Unmarshal(t.Ctx.Input.RequestBody, &task); err != nil || !task.IsValidTask() {
			t.Ctx.Output.SetStatus(http.StatusBadRequest)
			t.Data["json"] = map[string]interface{}{"err": "bad body request"}
		} else if err = task.Submit(t.Ctx); err != nil {
			t.Ctx.Output.SetStatus(http.StatusInternalServerError)
			t.Data["json"] = map[string]interface{}{"err": err.Error()}
		} else {
			t.Data["json"] = map[string]interface{}{"data": task}
		}
	}
	defer t.Ctx.Request.Body.Close()
	_ = t.ServeJSON()
}