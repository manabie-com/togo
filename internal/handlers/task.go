package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/services/tasks"
	"github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
)

func listTasks(service tasks.ToDoService) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		id, ok := userIDFromCtx(req.Context())
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
		tasks, err := service.ListTasks(req.Context(), id, utils.Value(req, "created_date"))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(resp).Encode(map[string][]*storages.Task{
			"data": tasks,
		})
	})
}

func addTask(service tasks.ToDoService) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		id, ok := userIDFromCtx(req.Context())
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
		t := &storages.Task{}
		err := json.NewDecoder(req.Body).Decode(t)
		defer req.Body.Close()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		task, err := service.AddTask(req.Context(), utils.ConvertStringToSqlNullString(id), *t)
		if err != nil {
			if err.Error() == "exceed today maximum allowed number of tasks" {
				resp.WriteHeader(http.StatusConflict)
			} else {
				resp.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(resp).Encode(map[string]*storages.Task{
			"data": task,
		})
	})
}

func deleteTaskByDate(service tasks.ToDoService) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")

		createdDate := utils.Value(req, "created_date")
		userID, _ := userIDFromCtx(req.Context())

		err := service.DeleteTaskByDate(
			req.Context(),
			utils.ConvertStringToSqlNullString(userID),
			createdDate,
		)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		resp.WriteHeader(http.StatusNoContent)
	})
}

type userAuthKey int8

func validToken(userService users.ToDoService) negroni.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		resp.Header().Set("Content-Type", "application/json")
		token := req.Header.Get("Authorization")

		claims := make(jwt.MapClaims)
		t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(userService.JWTKey), nil
		})
		if err != nil {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !t.Valid {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, ok := claims["user_id"].(string)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := req.Context()
		if !userService.FindUserById(ctx, utils.ConvertStringToSqlNullString(id)) {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, userAuthKey(0), id)
		next(resp, req.WithContext(ctx))
	}
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

func MakeTaskHandler(r *mux.Router, n negroni.Negroni, taskService tasks.ToDoService, userService users.ToDoService) {
	r.Handle("/tasks", n.With(
		validToken(userService),
		negroni.Wrap(listTasks(taskService)),
	)).Methods("GET").Name("list tasks")
	r.Handle("/tasks", n.With(
		validToken(userService),
		negroni.Wrap(addTask(taskService)),
	)).Methods("POST").Name("add task")
	r.Handle("/tasks", n.With(
		validToken(userService),
		negroni.Wrap(deleteTaskByDate(taskService)),
	)).Methods("DELETE").Name("delete task")
}
