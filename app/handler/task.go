package handler

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/huuthuan-nguyen/manabie/app/render"
	"github.com/huuthuan-nguyen/manabie/app/request"
	"github.com/huuthuan-nguyen/manabie/app/transformer"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"log"
	"net/http"
	"strconv"
)

// TaskStore /**
func (handler *Handler) TaskStore(w http.ResponseWriter, r *http.Request) {
	taskRequest := request.Task{}
	taskModel := model.Task{}

	if err := taskRequest.Bind(r, &taskModel); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := taskModel.Create(r.Context(), handler.db); err != nil {
		log.Println(err)
		if _, ok := err.(*model.ErrExceedDailyLimit); ok {
			utils.PanicBadRequest(err)
			return
		}
		utils.PanicInternalServerError(err)
		return
	}

	taskTransformer := &transformer.TaskTransformer{}
	taskItem := transformer.NewItem(taskModel, taskTransformer)

	// render task item to JSON
	render.JSON(w, r, taskItem)
}

// TaskUpdate /**
func (handler *Handler) TaskUpdate(w http.ResponseWriter, r *http.Request) {
	taskRequest := request.Task{}
	taskModel := model.Task{}
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	_, err = model.FindOneTaskByID(r.Context(), taskID, handler.db)
	if err != nil {
		utils.PanicNotFound()
		return
	}

	taskModel.ID = taskID
	if err := taskRequest.Bind(r, &taskModel); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := taskModel.Update(r.Context(), handler.db); err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	taskTransformer := &transformer.TaskTransformer{}
	taskItem := transformer.NewItem(taskModel, taskTransformer)

	// render task item to JSON
	render.JSON(w, r, taskItem)
}

// TaskDestroy /**
func (handler *Handler) TaskDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	_, err = model.FindOneTaskByID(r.Context(), taskID, handler.db)
	if err != nil {
		utils.PanicNotFound()
		return
	}

	taskModel := model.Task{}
	taskModel.ID = taskID

	if err := taskModel.Delete(r.Context(), handler.db); err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	render.NoContent(w, r)
}

// TaskIndex /**
func (handler *Handler) TaskIndex(w http.ResponseWriter, r *http.Request) {
	// prepare data response
	tasks, err := model.FindTasks(r.Context(), handler.db, r.URL.Query())
	if err != nil {
		render.Error(w, r, err)
		return
	}

	taskTransformer := &transformer.TaskTransformer{}
	taskCollection := transformer.NewCollection(tasks, taskTransformer)

	// render Collection to JSON
	render.JSON(w, r, taskCollection)
}

// TaskShow /**
func (handler *Handler) TaskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	task, err := model.FindOneTaskByID(r.Context(), taskID, handler.db)
	if err != nil {
		utils.PanicNotFound()
		return
	}

	taskTransformer := &transformer.TaskTransformer{}
	taskItem := transformer.NewItem(task, taskTransformer)

	// render task item to JSON
	render.JSON(w, r, taskItem)
}
