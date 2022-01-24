package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"togo/common"
)

type ServerContext interface {
	GetService(prefix string) (interface{})
}

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
			fmt.Printf("%v is running\n", sv.GetPrefix())
			stop <- sv.Run()
		}(svc)
	}

	go func() {
		if err := s.restHandler().Run(fmt.Sprintf(":%v", s.port)); err != nil {
			stop <- err
		}
	}()

	for  {
		select {
		case err := <- stop:
			if err != nil {
				return err
			}

		case sig := <-sigs:
			if sig != nil {
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
		return errors.New(common.PortNullErr)
	}

	if s.restHandler == nil {
		return errors.New(common.ServiceNullErr)
	}

	return nil
}

func (s *server) initFlags() error {
	return nil
}

func (s *server) InitService(svc service) {
	if has, ok := s.services[svc.GetPrefix()]; ok {
		log.Fatal(fmt.Sprintf("Service %v is duplicated", has.GetPrefix()))
	}

	s.services[svc.GetPrefix()] = svc
}

func (s *server) AddHandler(hdl func() *gin.Engine) {
	s.restHandler = hdl
}

func (s *server) GetService(prefix string) (interface{}) {
	if svc, ok := s.services[prefix]; ok {
		return svc.Get()
	}

	return nil
}