package usergrpc

import (
	"context"
	"encoding/json"

	goservice "github.com/phathdt/libs/go-sdk"
	protos "github.com/phathdt/libs/togo_proto/out/proto"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/user/userrepo"
	"togo/modules/user/userstorage"
)

type userGrpcServer struct {
	sc goservice.ServiceContext
}

func NewUserGrpcServer(sc goservice.ServiceContext) *userGrpcServer {
	return &userGrpcServer{sc: sc}
}

func (u *userGrpcServer) FindUser(ctx context.Context, request *protos.FindUserRequest) (*protos.FindUserResponse, error) {
	sc := u.sc

	var cond map[string]interface{}
	if err := json.Unmarshal([]byte(request.Cond), &cond); err != nil {
		panic(err)
	}
	db := sc.MustGet(common.DBMain).(*gorm.DB)
	store := userstorage.NewSQLStore(db)
	repo := userrepo.NewRepo(store)

	user, err := repo.FindUser(ctx, cond)
	if err != nil {
		panic(err)
	}

	return &protos.FindUserResponse{User: &protos.SimpleUser{
		Id:    int32(user.ID),
		Email: user.Email,
	}}, nil

}

func (u *userGrpcServer) GetUserLimit(ctx context.Context, request *protos.GetUserLimitRequest) (*protos.GetUserLimitResponse, error) {
	sc := u.sc
	db := sc.MustGet(common.DBMain).(*gorm.DB)
	store := userstorage.NewSQLStore(db)
	repo := userrepo.NewRepo(store)

	user, err := repo.FindUser(ctx, map[string]interface{}{"id": request.UserId})
	if err != nil {
		panic(err)
	}

	return &protos.GetUserLimitResponse{Limit: int32(user.LimitTask)}, nil
}
