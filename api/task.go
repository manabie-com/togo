package api

import (
	"net/http"
	"strconv"
	"togo/config"
	"togo/msg"
	"togo/services"
	"togo/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTask(c *gin.Context) {
	appG := Gin{C: c}
	offset, limit := utils.GetPage(c, config.GetConfig().DefaultPageNum, config.GetConfig().DefaultPageLimit)
	service := services.TaskReq{
		PageNum:  offset,
		PageSize: limit,
	}
	list, err := service.GetAll()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}
	total, err := service.GetTotal()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = list
	data["total"] = total
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), data, nil)
}

func AddTask(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	waitUse := claims.(*utils.CustomClaims)

	var service services.TaskReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.UserID = waitUse.Id
		objRes, err := service.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func UpdateTask(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var service services.TaskReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.ID = uint(id)
		objRes, err := service.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func DeleteTask(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var service services.TaskReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.ID = uint(id)
		objRes, err := service.Delete()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}
