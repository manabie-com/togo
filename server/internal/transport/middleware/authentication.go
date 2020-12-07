package middleware

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"strconv"
)

func Authenticate(endpoint kit.Endpoint) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logger.Debugf("[Middleware][Authentication] Start Authenticate: %v", httprouter.ParamsFromContext(ctx))

		auth := ctx.Value(define.ContextKeyAuthorization).(string)

		claims, err := utils.GetClaimsFromToken(auth, viper.GetString("jwt.key"))
		if err != nil {
			logger.Errorf("[Middleware][Authentication] getClaim error: %s", err.Error())
			return nil, define.Unauthenticated
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			logger.Error("[Middleware][Authentication]: userId not found")
			return nil, define.Unauthenticated
		}

		uint64Id, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			logger.Error("[Middleware][Authentication]: userId not parsed")
			return nil, define.Unauthenticated
		}

		ctx = context.WithValue(ctx, define.ContextKeyUserID, uint64Id)
		return endpoint(ctx, request)
	}
}