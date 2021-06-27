package routers

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/delivery/rest"
	"log"
	"net/http"
)

type Server struct {
	Serializer *rest.Serializer
	HTTP       config.HTTPConf
}

func NewServer(serializer *rest.Serializer, http config.HTTPConf) *Server {
	return &Server{
		Serializer: serializer,
		HTTP: http,
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
		req, valid := s.Serializer.ValidToken(resp, req)
		if !valid {
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
	http.ListenAndServe(s.HTTP.Addr, s)
}
