package transport

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/utils"
	"log"
	"net/http"
	"time"
)
type TaskTransport interface{
	ListTasks( resp http.ResponseWriter, req *http.Request)
	AddTasks( resp http.ResponseWriter, req *http.Request)
}
type taskTransport struct {
	taskService services.TaskService
}

func NewTaskTransport(ts services.TaskService) TaskTransport{
	return &taskTransport{
		taskService: ts,
	}
}


func (tt *taskTransport) ListTasks( resp http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	userId, _ := userIDFromCtx(ctx)
	createDate := req.FormValue("created_date")
	res, err := tt.taskService.ListTasks(ctx,userId, createDate)

	if err != nil {
		utils.HttpResponseInternalServerError(resp, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*model.Task{
		"data": res,
	})
}

func (tt *taskTransport) AddTasks( resp http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	userID, _ := userIDFromCtx(ctx)

	if !tt.taskService.IsAllowedToAddTask(ctx, userID) {
		utils.HttpResponseBadRequest(resp, "You have added more than the number of tasks allowed per day!")
		return
	}

	t := &model.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		utils.HttpResponseInternalServerError(resp, err.Error())
		return
	}

	now := time.Now()
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")
	err = tt.taskService.AddTask(ctx, t)

	if err != nil {
		utils.HttpResponseInternalServerError(resp, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string]*model.Task{
		"data": t,
	})
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value("userId")
	id, ok := v.(string)
	return id, ok
}

func ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)

	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte("wqGyEBBfPK9w3Lxw"), nil
	})

	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)

	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), "userId", id))
	return req, true
}
