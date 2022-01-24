package service

import (
	"context"
	"session_service/proto"
)

//serviceServer ...
type serviceServer struct {
	proto.UnimplementedSessionServiceServer
}

//NewSessionServiceServer ...
func NewSessionServiceServer() proto.SessionServiceServer {
	return &serviceServer{}
}

func (s *serviceServer) GetAccountIDFromToken(ctx context.Context, request *proto.TokenString) (*proto.AccountID, error) {
	return nil, nil
}

func (s *serviceServer) GetAccountTypeFromToken(ctx context.Context, request *proto.TokenString) (*proto.AccountType, error) {
	return nil, nil
}

func (s *serviceServer) CreateToken(ctx context.Context, request *proto.AccountInfo) (*proto.TokenString, error) {
	return nil, nil
}

func (s *serviceServer) RefreshToken(ctx context.Context, request *proto.TokenString) (*proto.TokenString, error) {
	return nil, nil
}

func (s *serviceServer) DeleteToken(ctx context.Context, request *proto.TokenString) (*proto.Status, error) {
	return nil, nil
}

func (s *serviceServer) CheckToken(ctx context.Context, request *proto.TokenString) (*proto.Status, error) {
	return nil, nil
}
