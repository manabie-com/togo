package main

import (
	"fmt"
	"main/config"
	"main/internal/service"
	"net/http"
)

type Server struct {
	http.Server
	cfg     config.Config
	service service.Service
}

func NewServer(cfg config.Config, svc service.Service) Server {
	return Server{
		cfg:     cfg,
		service: svc,
	}
}

func (s *Server) start() error {
	muxHttp := http.NewServeMux()
	muxHttp.Handle("/api/v1/todo", s.service.CreateTodoHandler())
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.HTTPAddress), muxHttp)
}
