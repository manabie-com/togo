package handle

import (
	"ManabieProject/helper"
	"ManabieProject/src/model/requestmodel"
	"ManabieProject/src/service/process"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAccount Register an account
func RegisterAccount(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var registerRequestBody requestmodel.RegisterRequest
	err := decoder.Decode(&registerRequestBody)
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

	ok, response := process.RegisterAccountProcess(&registerRequestBody)
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

// LoginAccount account login
func LoginAccount(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var loginRequestBody requestmodel.LoginRequest
	err := decoder.Decode(&loginRequestBody)
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

	ok, response := process.LoginAccountProcess(&loginRequestBody)
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
