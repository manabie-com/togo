package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"go/types"
	"io/ioutil"
	"time"
	"togo/domain"
	"togo/interface/controllers/structs"
	"togo/usecase"
)

type TodoController struct {
	todoUsecase      usecase.TodoUsecaseInterface
	todoLimitUsecase usecase.TodoLimitUsecaseInterface
}

func NewTodoController(us usecase.TodoUsecaseInterface, usLimit usecase.TodoLimitUsecaseInterface) TodoControllerInterface {
	return &TodoController{us, usLimit}
}

func (receiver TodoController) CreateTodoUser(c *gin.Context) {
	var params = make(map[string]interface{})

	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code":    "error",
			"message": "Error read param API",
			"data":    types.Interface{},
		})
		return
	}

	err = json.Unmarshal(jsonBytes, &params)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code":    "error",
			"message": "unmarshal body: " + err.Error(),
			"data":    types.Interface{},
		})
		return
	}

	var createStructCtl structs.CreateStruct
	var createResultStructCtl structs.CreateResultStruct

	// TODO check from cache

	// Check from db
	var getStructCtl structs.GetStruct
	var getResultStructCtl structs.GetResultStruct
	getStructCtl.Conditions = fmt.Sprintf("user_id = %s AND created_at >= '%s'", params["user_id"], time.Now().Format("2006-01-02"))

	getStructUsecase := getStructCtl.ConvertGetInterfaceToUsecase()
	getResultStructUsecase := receiver.todoUsecase.Get(getStructUsecase)
	getResultStructCtl = getResultStructCtl.ConvertGetResultUsecaseToInterface(getResultStructUsecase)

	numOfTodo := getResultStructCtl.Data.RowsAffected

	getStructCtl.Conditions = map[string]interface{}{
		"user_id": params["user_id"],
	}
	getStructUsecase = getStructCtl.ConvertGetInterfaceToUsecase()
	getResultStructUsecase = receiver.todoLimitUsecase.Get(getStructUsecase)
	getResultStructCtl = getResultStructCtl.ConvertGetResultUsecaseToInterface(getResultStructUsecase)

	var todoLimit []domain.TodoLimit

	todoLimitRow := getResultStructCtl.Data.Result

	todoLimitByte, err := json.Marshal(todoLimitRow)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code":    "error",
			"message": "Marshal: " + err.Error(),
			"data":    types.Interface{},
		})
		return
	}

	err = json.Unmarshal(todoLimitByte, &todoLimit)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code":    "error",
			"message": "unmarshal: " + err.Error(),
			"data":    types.Interface{},
		})
		return
	}

	firstRowIndex := 0

	if numOfTodo >= int64(todoLimit[firstRowIndex].Limit) {
		c.JSON(200, map[string]interface{}{
			"code":    "error",
			"message": "You have reach limit of todo today",
			"data":    types.Interface{},
		})
		return
	}

	// TODO Valid input
	createStructCtl.Data = map[string]interface{}{
		"id":          uuid.New().String(),
		"user_id":     params["user_id"],
		"description": params["description"],
		"status":      "new",
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	createStructUsecase := createStructCtl.ConvertCreateInterfaceToUsecase()
	createResultStructUsecase := receiver.todoUsecase.Create(createStructUsecase)
	createResultStructCtl = createResultStructCtl.ConvertCreateResultUsecaseToInterface(createResultStructUsecase)

	c.JSON(200, map[string]interface{}{
		"code":    createResultStructCtl.Status,
		"message": createResultStructCtl.Message,
		"data":    createResultStructCtl.Data,
	})
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
		"data": map[string]interface{}{
			"total": getResultStructCtl.Data.RowsAffected,
			"list":  getResultStructCtl.Data.Result,
		},
	})
}
