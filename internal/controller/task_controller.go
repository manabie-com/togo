package controller

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/response"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/services/impl"
	"github.com/manabie-com/togo/internal/utils"
	"io/ioutil"
	"net/http"
)

type TaskController struct {
	service services.TasksService
}

func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{impl.NewTaskServiceImpl(db) }
}

func (s *TaskController) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := utils.UerIDFromCtx(req.Context())
	tasks, err := s.service.RetrieveTasks(id, value(req, "created_date"))

	if err != nil {
		response.RespondWithError(resp, http.StatusInternalServerError, err.Error())
	} else {
		response.RespondWithJSON(resp, http.StatusOK, tasks)
	}

}


func (s *TaskController) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &model.Task{}
	userID, _ := utils.UerIDFromCtx(req.Context())
	// write body
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.RespondWithError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	errPartJson := json.Unmarshal(data, t)
	if errPartJson != nil {
		response.RespondWithError(resp, http.StatusInternalServerError, errPartJson.Error())
		return
	}
	defer req.Body.Close()

	t, errDb := s.service.AddTask(t, userID)
	if errDb != nil {
		response.RespondWithError(resp, http.StatusInternalServerError, errDb.Error())
		return
	}
	response.RespondWithJSON(resp, http.StatusOK, t)
}