package service

import (
	"context"

	"session_service/auth"
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

	td, err := auth.GenerateToken(request)
	if err != nil {
		return nil, err
	}

	err = auth.CreateAuth(request.GetId(), td)
	if err != nil {
		return nil, err
	}

	return &proto.TokenString{
		Token: td.AccessToken,
	}, nil
}

func (s *serviceServer) RefreshToken(ctx context.Context, request *proto.TokenString) (*proto.TokenString, error) {

	//Delete old Access token
	tokenMetadata, err := auth.ExtractTokenMetadata(request.GetToken())
	if err != nil {
		return nil, err
	}

	deleted, delErr := auth.DeleteAuth(tokenMetadata.AccessUUID)
	if delErr != nil || deleted == 0 {
		return nil, delErr
	}

	//Create new Access token
	td, createErr := auth.GenerateToken(&proto.AccountInfo{
		Id:   tokenMetadata.AccountID,
		Type: tokenMetadata.AccountType,
	})
	if createErr != nil {
		return nil, createErr
	}

	//Save token Metadata to Redis
	err = auth.CreateAuth(tokenMetadata.AccountID, td)
	if err != nil {
		return nil, err
	}

	return &proto.TokenString{
		Token: td.AccessToken,
	}, nil
}

func (s *serviceServer) DeleteToken(ctx context.Context, request *proto.TokenString) (*proto.Status, error) {

	tokenMetadata, err := auth.ExtractTokenMetadata(request.GetToken())
	if err != nil {
		return nil, err
	}

	deleted, delErr := auth.DeleteAuth(tokenMetadata.AccessUUID)
	if delErr != nil || deleted == 0 {
		return nil, delErr
	}

	return &proto.Status{
		Success: true,
	}, nil
}

func (s *serviceServer) CheckToken(ctx context.Context, request *proto.TokenString) (*proto.Status, error) {

	tokenMetadata, err := auth.ExtractTokenMetadata(request.GetToken())
	if err != nil {
		return nil, err
	}

	_, authErr := auth.FetchAuth(tokenMetadata)
	if authErr != nil {
		return nil, authErr
	}

	return &proto.Status{
		Success: true,
	}, nil
}
