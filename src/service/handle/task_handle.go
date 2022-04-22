package handle

import (
	"ManabieProject/helper"
	"ManabieProject/src/model/requestmodel"
	"ManabieProject/src/service/process"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTask create new task
func CreateTask(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var createTaskRequestBody requestmodel.CreateTaskRequest
	err := decoder.Decode(&createTaskRequestBody)
	if err != nil {
		response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Incorrect data format !"}
		responseByte, err := json.Marshal(response)
		if err != nil {
			ctx.Data(response.Status, "application/json", nil)
		} else {
			ctx.Data(response.Status, "application/json", responseByte)
		}
		return
	}
	ok, response := process.CreateTaskProcess(&createTaskRequestBody)
	if !ok {
		responseByte, _ := json.Marshal(response.(helper.Problem))
		if err != nil {
			ctx.Data(response.(helper.Problem).Status, "application/json", nil)
		} else {
			ctx.Data(response.(helper.Problem).Status, "application/json", responseByte)
		}
		return
	}
	responseByte, _ := json.Marshal(response.(helper.Success))
	ctx.Data(response.(helper.Success).Status, "application/json", responseByte)
	return
}

// UpdateTask update task
func UpdateTask(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var updateTaskRequestBody requestmodel.UpdateTaskRequest
	err := decoder.Decode(&updateTaskRequestBody)
	if err != nil {
		response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Incorrect data format !"}
		responseByte, err := json.Marshal(response)
		if err != nil {
			ctx.Data(response.Status, "application/json", nil)
		} else {
			ctx.Data(response.Status, "application/json", responseByte)
		}
		return
	}
	ok, response := process.UpdateTaskProcess(&updateTaskRequestBody)
	if !ok {
		responseByte, _ := json.Marshal(response.(helper.Problem))
		if err != nil {
			ctx.Data(response.(helper.Problem).Status, "application/json", nil)
		} else {
			ctx.Data(response.(helper.Problem).Status, "application/json", responseByte)
		}
		return
	}
	responseByte, _ := json.Marshal(response.(helper.Success))
	ctx.Data(response.(helper.Success).Status, "application/json", responseByte)
	return
}
