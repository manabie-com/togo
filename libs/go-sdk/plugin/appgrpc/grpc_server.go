package appgrpc

import (
	"fmt"
	"net"
	"time"

	"flag"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/phathdt/libs/go-sdk/logger"
	"github.com/phathdt/libs/go-sdk/plugin/appgrpc/middleware/logging"
	"github.com/phathdt/libs/go-sdk/plugin/appgrpc/middleware/recovermiddleware"
	"google.golang.org/grpc"
)

type gprcServer struct {
	prefix      string
	port        int
	server      *grpc.Server
	logger      logger.Logger
	registerHdl func(*grpc.Server)
}

func NewGRPCServer(prefix string) *gprcServer {
	return &gprcServer{prefix: prefix}
}

func (s *gprcServer) SetRegisterHdl(hdl func(*grpc.Server)) {
	s.registerHdl = hdl
}

func (s *gprcServer) GetPrefix() string {
	return s.prefix
}

func (s *gprcServer) Get() interface{} {
	return s
}

func (s *gprcServer) Name() string {
	return s.prefix
}

func (s *gprcServer) InitFlags() {
	flag.IntVar(&s.port, s.GetPrefix()+"-port", 50051, "Port of gRPC service")
}

func (s *gprcServer) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.prefix)

	s.logger.Infoln("Setup gRPC service:", s.prefix)
	s.server = grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			recovermiddleware.StreamServerInterceptor(),
			logging.StreamServerInterceptor(s.logger),
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			recovermiddleware.UnaryServerInterceptor(),
			logging.UnaryServerInterceptor(s.logger),
		)),
	)

	return nil
}

func (s *gprcServer) Recover() {
	if err := recover(); err != nil {
		s.logger.Error("recover error", err)
	}
}

func (s *gprcServer) Run() error {
	go func() {
		defer s.Recover()

		time.Sleep(time.Second * 5)

		_ = s.Configure()

		if s.registerHdl != nil {
			s.logger.Infoln("registering services...")
			s.registerHdl(s.server)
		}

		address := fmt.Sprintf("0.0.0.0:%d", s.port)
		lis, err := net.Listen("tcp", address)

		if err != nil {
			s.logger.Errorln("Error %v", err)
		}

		s.server.Serve(lis)
	}()

	return nil
}

func (s *gprcServer) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		s.server.Stop()
		c <- true
		s.logger.Infoln("Stopped")
	}()
	return c
}
