package middleware

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/kit"
)

func Authenticate(endpoint kit.Endpoint) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}

//func validToken(token, jwtKey string) error {
//	claims := make(jwt.MapClaims)
//	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
//		return []byte(jwtKey), nil
//	})
//	if err != nil {
//		return err
//	}
//
//	if !t.Valid {
//		return define.Unauthenticated
//	}
//
//	_, ok := claims["user_id"].(string)
//	if !ok {
//		return define.Unauthenticated
//	}
//
//	return nil
//}