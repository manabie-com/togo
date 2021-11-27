package main

import (
	"fmt"
	"main/config"
	"main/internal/service"
	"net/http"
)

type Server struct {
	http.Server
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) start(service *service.TogoService) error {
	muxHttp := http.NewServeMux()

	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.HTTPAddress), muxHttp)
}
