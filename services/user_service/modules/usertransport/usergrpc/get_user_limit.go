package usergrpc

import (
	"context"

	protos "github.com/phathdt/libs/togo_proto/out/proto"
	"gorm.io/gorm"
	"user_service/common"
	"user_service/modules/userstorage"
)

func (s *userGrpcServer) GetUserLimit(ctx context.Context, request *protos.GetUserLimitRequest) (*protos.GetUserLimitResponse, error) {
	sc := s.sc
	db := sc.MustGet(common.DBMain).(*gorm.DB)

	store := userstorage.NewSQLStore(db)

	user, err := store.FindUser(ctx, map[string]interface{}{"id": request.UserId})
	if err != nil {
		panic(err)
	}

	return &protos.GetUserLimitResponse{Limit: int32(user.LimitTask)}, nil
}
