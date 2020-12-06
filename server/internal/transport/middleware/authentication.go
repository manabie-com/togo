package middleware

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func Authenticate(endpoint kit.Endpoint) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token := httprouter.ParamsFromContext(ctx).ByName("Authorization")
		claims, err := utils.GetClaimsFromToken(token, viper.GetString("jwt.key"))
		if err != nil {
			return nil, define.Unauthenticated
		}
		id, ok := claims["user_id"].(string)
		if !ok {
			return nil, define.Unauthenticated
		}

		ctx = context.WithValue(ctx, define.ContextKeyUserID, id)
		return endpoint(ctx, request)
	}
}