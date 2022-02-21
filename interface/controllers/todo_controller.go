package controllers

import (
	"github.com/gin-gonic/gin"
	"togo/interface/controllers/structs"
	"togo/usecase"
)

type TodoController struct {
	todoUsecase usecase.TodoUsecaseInterface
}

func NewTodoController(us usecase.TodoUsecaseInterface) TodoControllerInterface {
	return &TodoController{us}
}

func (receiver TodoController) GetAllTodoUser(c *gin.Context) {

	userId := c.Query("user_id")

	var getStructCtl structs.GetStruct
	var getResultStructCtl structs.GetResultStruct

	getStructCtl.Conditions = map[string]interface{}{
		"user_id": userId,
	}

	getStructUsecase := getStructCtl.ConvertGetInterfaceToUsecase()
	getResultStructUsecase := receiver.todoUsecase.Get(getStructUsecase)
	getResultStructCtl = getResultStructCtl.ConvertGetResultUsecaseToInterface(getResultStructUsecase)

	c.JSON(200, map[string]interface{}{
		"code":    getResultStructCtl.Status,
		"message": getResultStructCtl.Message,
		"data":    getResultStructCtl.Data,
	})
}
