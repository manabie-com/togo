package service

import (
	"github.com/tonghia/togo/internal/store"
	"github.com/tonghia/togo/pkg/tasklimit"
)

type Service struct {
	store        store.Querier
	userLimitSvc tasklimit.UserLimitSvc
}

func NewService(store store.Querier, userLimitSvc tasklimit.UserLimitSvc) *Service {
	return &Service{store, userLimitSvc}
}
