package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase"
)

const (
	LAYOUT = "2006-01-02"
)

type TaskHandler struct {
	TUsecase usecase.TaskUsecase
}

func NewTaskHandler(us usecase.TaskUsecase) TaskHandler {
	return TaskHandler{
		TUsecase: us,
	}
}

func (t *TaskHandler) ListTasks(resp http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	user_id, _ := userIDFromCtx(ctx)
	created_date := req.FormValue("created_date")
	created_dateT, _ := time.Parse(LAYOUT, created_date)
	tasks, err := t.TUsecase.ListTasks(ctx, user_id, created_dateT)

	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	json.NewEncoder(resp).Encode(map[string][]storages.Task{
		"data": tasks,
	})
	return nil
}

func (t *TaskHandler) AddTask(resp http.ResponseWriter, req *http.Request) error {
	var task storages.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return err
	}
	userID, _ := userIDFromCtx(req.Context())
	ctx := req.Context()
	numTasks, err := t.TUsecase.CountTaskPerDay(ctx, userID, time.Now())
	if numTasks > 5 {
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Number of task exceed",
		})
		return errors.New("Exceed number of task today")
	}
	if err != nil {
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	task.UserID = userID
	task.CreatedDate = time.Now()
	err = t.TUsecase.AddTask(req.Context(), &task)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": &task,
	})
	return nil
}
func (t *TaskHandler) GetAuthToken(resp http.ResponseWriter, req *http.Request) error {
	id := req.FormValue("user_id")
	pass := req.FormValue("password")

	val, err := t.TUsecase.ValidateUser(req.Context(), id, pass)
	if err != nil || !val {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return err
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := CreateToken(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
	return nil
}

func (t *TaskHandler) ValidateToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	tok, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !tok.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
