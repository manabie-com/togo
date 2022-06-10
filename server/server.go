package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"togo/common"
	logger2 "togo/pkg/logger"
)

type service interface {
	Run() error
	Stop() <-chan bool
	GetPrefix() string
	Get() interface{}
}

type server struct {
	prefix      string
	port        uint
	services    map[string]service
	restHandler func() *gin.Engine
	logger      logger2.Loggers
}

func NewServer(prefix string, port uint) *server {
	svs := make(map[string]service)
	return &server{prefix: prefix, port: port, services: svs}
}

func (s *server) Run() error {
	if err := s.configure(); err != nil {
		return err
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	stop := make(chan error, 1)

	for _, svc := range s.services {
		go func(sv service) {
			if err := sv.Run(); err != nil {
				stop <- err
			} else {
				s.logger.Info().Println(fmt.Sprintf("%v is running", sv.GetPrefix()))
			}

		}(svc)
	}

	go func() {
		if err := s.restHandler().Run(fmt.Sprintf(":%v", s.port)); err != nil {
			stop <- err
		}
	}()

	for {
		select {
		case err := <-stop:
			if err != nil {
				return err
			}

		case sig := <-sigs:
			if sig != nil {
				for _, svc := range s.services {
					svc.Stop()
					s.logger.Info().Println(fmt.Sprintf("%v is stopped", svc.GetPrefix()))
				}

				return errors.New(sig.String())
			}
		}
	}

	return nil
}

func (s *server) configure() error {
	if err := s.initFlags(); err != nil {
		return err
	}

	if s.port == 0 {
		return errors.New(common.DataIsNullErr("Port"))
	}

	if s.restHandler == nil {
		return errors.New(common.DataIsNullErr("RestAPI"))
	}

	return nil
}

func (s *server) initFlags() error {
	return nil
}

func (s *server) InitService(svc service) {
	if has, ok := s.services[svc.GetPrefix()]; ok {
		s.logger.Error().Fatal(fmt.Sprintf("Service %v is duplicated", has.GetPrefix()))
	}

	s.services[svc.GetPrefix()] = svc
}

func (s *server) AddHandler(hdl func() *gin.Engine) {
	s.restHandler = hdl
}

func (s *server) AddLogger(loggers logger2.Loggers) {
	s.logger = loggers
}

func (s *server) GetService(prefix string) interface{} {
	if svc, ok := s.services[prefix]; ok {
		return svc.Get()
	}

	return nil
}

func (s *server) GetLogger() logger2.Loggers {
	return s.logger
}
