package routers

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/delivery/rest"
	"github.com/manabie-com/togo/internal/utils"
	"log"
	"net/http"
)

type Server struct {
	Serializer *rest.Serializer
}

func NewServer(serializer *rest.Serializer) *Server {
	return &Server{
		Serializer: serializer,
	}
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.Serializer.GetAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.Serializer.ValidToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(resp).Encode(
				utils.BuildErrorResponseRequest(&utils.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				}))
			return
		}
		switch req.Method {
		case http.MethodGet:
			s.Serializer.ListTasks(resp, req)
		case http.MethodPost:
			s.Serializer.AddTask(resp, req)
		}
		return
	}
}

func (s *Server) Run() {
	http.ListenAndServe(":5050", s )
}
