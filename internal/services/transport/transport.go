package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/services/usecase"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/libs"
)

type transport struct {
	useCase storages.EntityUseCase
}

func Init(db *gorm.DB, method string, req *http.Request) (res map[string]interface{}, err error) {
	repository := postgres.InitRepository(db)
	useCase := usecase.InitUseCase(repository)
	transport := transport{
		useCase: useCase,
	}

	switch method {
	case "get-auth":
		res, err = transport.GetAuthToken(req)
	case "list-tasks":
		res, err = transport.ListTasks(req)
	case "add-task":
		res, err = transport.AddTask(req)
	}

	return res, err
}

func (trans *transport) GetAuthToken(req *http.Request) (res map[string]interface{}, err error) {

	var (
		args = map[string]string{}
		user = libs.Value(req, "user_id")
		pwd  = libs.Value(req, "password")
	)

	args["user_id"] = user.String
	args["password"] = pwd.String

	res, err = trans.useCase.GetAuthToken(args)
	return res, err
}

func (trans *transport) ListTasks(req *http.Request) (res map[string]interface{}, err error) {

	var (
		args       = map[string]string{}
		createDate = libs.Value(req, "created_date")
	)

	args["user_id"], _ = libs.UserIDFromCtx(req.Context())
	args["created_date"] = createDate.String

	res, err = trans.useCase.ListTasks(args)
	return res, err
}

func (trans *transport) AddTask(req *http.Request) (res map[string]interface{}, err error) {

	var (
		task = storages.Task{}
		now  = time.Now()
	)

	err = json.NewDecoder(req.Body).Decode(&task)
	defer req.Body.Close()

	task.UserID, _ = libs.UserIDFromCtx(req.Context())
	task.ID = uuid.New().String()
	task.CreatedDate = now.Format("2006-01-02")

	res, err = trans.useCase.AddTask(task)
	return res, err
}
