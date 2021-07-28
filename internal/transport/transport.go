package transport

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/internal/storages"
	taskStorage "github.com/manabie-com/togo/internal/storages/task"
	userStorage "github.com/manabie-com/togo/internal/storages/user"
	"github.com/manabie-com/togo/internal/usecase/task"
	"github.com/manabie-com/togo/internal/usecase/user"

	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
)

type Transport struct {
	TaskUsecase task.TaskUsecaseInterface
	UserUsecase user.UserUsecaseInterface
}

func NewTransport(db *gorm.DB) Transport {
	userS := userStorage.NewUserStorage(db)
	tranS := taskStorage.NewTaskStorage(db)
	return Transport{
		TaskUsecase: task.NewTaskUsecase(tranS),
		UserUsecase: user.NewUserUsecase(userS, tranS),
	}
}

func (t *Transport) Login(w http.ResponseWriter, r *http.Request) {
	params := LoginRequest{}
	if err := getFormParams(r, &params); err != nil {
		renderUnprocessibleEntityError(w, err.Error())
		return
	}

	if err := t.UserUsecase.ValidateUser(params.ID, params.Password); err != nil {
		renderUnauthorizedError(w, err.Error())
		return
	}

	respData, err := populateLoginResponse(storages.User{ID: params.ID})
	if err != nil {
		renderBadRequest(w, err.Error())
		return
	}

	renderSuccess(w, respData)
}

func (t *Transport) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := parseCreateTaskParams(r)
	if err != nil {
		renderBadRequest(w, err.Error())
		return
	}

	if err = t.UserUsecase.CreateTask(&task); err != nil {
		renderBadRequest(w, err.Error())
		return
	}

	renderSuccess(w, task)
}

func (t *Transport) ListTasks(w http.ResponseWriter, r *http.Request) {
	createdDate := r.FormValue("created_date")
	userID, ok := userIDFromCtx(r.Context())
	if !ok {
		renderBadRequest(w, "user_id is empty")
		return
	}

	tasks, err := t.TaskUsecase.RetrieveTasks(userID, createdDate)
	if err != nil {
		renderBadRequest(w, err.Error())
		return
	}

	renderSuccess(w, tasks)
}

func getFormParams(req *http.Request, params interface{}) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	decoder := form.NewDecoder()
	if err := decoder.Decode(params, req.Form); err != nil {
		return err
	}

	return nil
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(0)
	id, ok := v.(string)
	return id, ok
}
