package api

import (
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
	"log"
	"net/http"
)

type IResponse interface {
	ToRes() interface{}
}

type TodoApi struct {
	author       AuthorApi
	task         TaskApi
	quotaService services.IQuotaService
}

func (s *TodoApi) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}
	ctx := req.Context()
	var err error
	var res IResponse
	switch req.URL.Path {
	case "/login":
		res, err = s.author.Login(ctx, req)
	case "/tasks":
		ctx, err = s.author.Validate(req)
		if err == nil {
			switch req.Method {
			case http.MethodGet:
				res, err = s.task.ListTasksByUserAndDate(ctx, req)
			case http.MethodPost:
				err = s.quotaService.LimitTask(ctx)
				if err == nil {
					res, err = s.task.AddTask(ctx, req)
				}
			}
		}
	}
	if err != nil {
		errToDo, ok := err.(*tools.TodoError)
		if !ok {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "Something when wrong",
			})
			return
		}
		resp.WriteHeader(errToDo.Code)
		json.NewEncoder(resp).Encode(errToDo.ToRes())
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(res.ToRes())
	return
}

func NewToDoApi(jwtKey string, db *sql.DB) TodoApi {
	return TodoApi{
		author:       NewAuthorApi(db, jwtKey),
		task:         NewTaskApi(db),
		quotaService: services.NewQuotaService(repos.NewQuotaRepo(db)),
	}
}
