package grpc

import (
	"github.com/khangjig/togo/proto"
	"github.com/khangjig/togo/usecase"
)

type TogoService struct {
	proto.UnimplementedTogoServiceServer
	UseCase *usecase.UseCase
}
