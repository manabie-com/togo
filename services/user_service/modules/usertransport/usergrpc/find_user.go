package usergrpc

import (
	"context"
	"encoding/json"

	protos "github.com/phathdt/libs/togo_proto/out/proto"
	"gorm.io/gorm"
	"user_service/common"
	"user_service/modules/userstorage"
)

func (s *userGrpcServer) FindUser(ctx context.Context, request *protos.FindUserRequest) (*protos.FindUserResponse, error) {
	cond := make(map[string]interface{})
	if err := json.Unmarshal([]byte(request.Cond), &cond); err != nil {
		panic(err)
	}

	sc := s.sc

	db := sc.MustGet(common.DBMain).(*gorm.DB)

	store := userstorage.NewSQLStore(db)

	user, err := store.FindUser(ctx, cond)
	if err != nil {
		panic(err)
	}

	return &protos.FindUserResponse{User: &protos.SimpleUser{
		Id:    int32(user.ID),
		Email: user.Email,
	}}, nil
}
