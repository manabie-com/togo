package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/storages/repo"
)

// TransportService implement HTTP server
type TransportService struct {
	JWTKey string
	DB     *sql.DB
}

func (this *TransportService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	u := &UserService{
		Common: this,
		UserStore: &repo.UserStore{
			DB: this.DB,
		},
	}
	t := &TaskService{
		TaskStore: &repo.TaskStore{
			DB: this.DB,
		},
		UserStore: &repo.UserStore{
			DB: this.DB,
		},
	}
	switch req.URL.Path {
	case "/signup":
		if req.Method != http.MethodPost {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		u.CreateUser(resp, req)
		return
	case "/login":
		if req.Method != http.MethodPost {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		u.GetAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = u.IsValidToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			t.ListTasks(resp, req)
		case http.MethodPost:
			t.CreateTask(resp, req)
		}
		return
	}
}
