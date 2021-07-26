package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/services/task"
	"github.com/manabie-com/togo/internal/services/user"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
)

func listTasks(service task.ToDoService) http.Handler {
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

type userAuthKey int8

func validToken(taskService task.ToDoService, userService user.ToDoService) negroni.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		resp.Header().Set("Content-Type", "application/json")
		token := req.Header.Get("Authorization")

		claims := make(jwt.MapClaims)
		t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(taskService.JWTKey), nil
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

func MakeTaskHandler(r *mux.Router, n negroni.Negroni, taskService task.ToDoService, userService user.ToDoService) {
	r.Handle("/tasks", n.With(
		validToken(taskService, userService),
		negroni.Wrap(listTasks(taskService)),
	)).Methods("GET", "OPTIONS").Name("list tasks")
}
