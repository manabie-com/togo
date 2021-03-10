package api

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/core"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/services"
	"net/http"
	"time"
)

type TodoCtrl struct {
	JWTKey      string
	TodoService *services.ToDoService
}

type userAuthKey int8

func (todoCtrl *TodoCtrl) registerRouter(router *mux.Router) *mux.Router {

	router.Path("/login").HandlerFunc(todoCtrl.getAuthToken)
	router.Methods(http.MethodGet).Path("/tasks").HandlerFunc(todoCtrl.listTasks)
	router.Methods(http.MethodPost).Path("/tasks").HandlerFunc(todoCtrl.addTask)

	return router
}

func (todoCtrl *TodoCtrl) listTasks(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	tasks, err := todoCtrl.TodoService.ListTasks(
		req.Context(),
		userID,
		getValueString(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*entities.Task{
		"data": tasks,
	})
}

func (todoCtrl *TodoCtrl) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &entities.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := userIDFromCtx(req.Context())

	resp.Header().Set("Content-Type", "application/json")

	err = todoCtrl.TodoService.AddTask(req.Context(), userID, t)
	if err != nil {
		if internalError, ok := err.(*core.InternalError); ok {
			if internalError.ErrCode == core.ERROR_CODE_EXCEED_TASK_LIMITS {
				resp.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(resp).Encode(map[string]interface{}{
					"error": err.Error(),
					"code":  internalError.ErrCode,
				})
			}
		} else {

			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
		}
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entities.Task{
		"data": t,
	})
}

func (todoCtrl *TodoCtrl) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := getValueString(req, "user_id")
	password := getValueString(req, "password")
	if !todoCtrl.TodoService.ValidateUser(req.Context(), id, password) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := todoCtrl.createToken(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	_ = json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

func (todoCtrl *TodoCtrl) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(todoCtrl.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func getValueString(req *http.Request, key string) string {
	return req.FormValue(key)
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
