package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"github.com/vchitai/togo/internal/store"
	"github.com/vchitai/togo/pb"
	"google.golang.org/grpc"
)

var (
	ll = l.New()
)

type serverImpl struct {
	cfg       *configs.Config
	isReady   bool
	toDoStore store.ToDo
}

func New(cfg *configs.Config, toDoStore store.ToDo) *serverImpl {
	return &serverImpl{
		cfg:       cfg,
		isReady:   true,
		toDoStore: toDoStore,
	}
}

func (s *serverImpl) RegisterWithServer(srv *grpc.Server) {
	pb.RegisterBaseServiceServer(srv, s)
	pb.RegisterToDoServiceServer(srv, s)
}

func (s *serverImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := pb.RegisterBaseServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := pb.RegisterToDoServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	return nil
}

// Close ...
func (s *serverImpl) Close(ctx context.Context) {
}
