package usergrpc

import (
	goservice "github.com/phathdt/libs/go-sdk"
)

type userGrpcServer struct {
	sc goservice.ServiceContext
}

func NewUserGrpcServer(sc goservice.ServiceContext) *userGrpcServer {
	return &userGrpcServer{sc: sc}
}
