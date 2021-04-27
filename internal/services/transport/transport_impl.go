package transport

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/services/usecase"
	"github.com/manabie-com/togo/internal/utils"
)

type TransportImpl struct {
	Usecase usecase.Usecase
	DB      *sql.DB
}

func NewTransport() *TransportImpl {
	db := models.Connect()
	return &TransportImpl{
		Usecase: usecase.NewUsecase(db),
		DB:      db,
	}
}

func (tr *TransportImpl) Login(resp http.ResponseWriter, req *http.Request) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		utils.ERROR(resp, http.StatusBadRequest, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		utils.ERROR(resp, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := tr.Usecase.GetAuthToken(req.Context(), user.Username, user.Password)
	if err != nil {
		utils.ERROR(resp, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(resp, http.StatusOK, token)
}

func (tr *TransportImpl) ListTasks(resp http.ResponseWriter, req *http.Request) {
	task := models.Task{}
	userId, ok := utils.UserIDFromCtx(req.Context())
	log.Println(userId)
	if !ok {
		utils.ERROR(resp, http.StatusNotFound, errors.New("user ID is not found"))
		return
	}

	createdDate := req.FormValue("created_date")
	task.UserID = userId
	task.CreatedDate = createdDate
	task.Prepare()
	err := task.Validate("retrieve_tasks")
	if err != nil {
		utils.ERROR(resp, http.StatusUnprocessableEntity, err)
		return
	}

	tasks, err := tr.Usecase.RetrieveTasks(req.Context(), userId, createdDate)
	if err != nil {
		utils.ERROR(resp, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(resp, http.StatusOK, tasks)
}

func (tr *TransportImpl) AddTask(resp http.ResponseWriter, req *http.Request) {
	task := &models.Task{}
	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		utils.ERROR(resp, http.StatusBadRequest, err)
		return
	}

	userId, ok := utils.UserIDFromCtx(req.Context())
	if !ok {
		utils.ERROR(resp, http.StatusNotFound, errors.New("user ID is not found"))
		return
	}

	task.Prepare()
	err = task.Validate("add_task")
	if err != nil {
		utils.ERROR(resp, http.StatusUnprocessableEntity, err)
		return
	}

	task, err = tr.Usecase.AddTask(req.Context(), userId, task)
	if err != nil {
		utils.ERROR(resp, http.StatusBadRequest, err)
		return
	}

	utils.JSON(resp, http.StatusCreated, task)
}
