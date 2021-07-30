package config

import (
	"togo/db/postgres"
)

type ServiceContext struct {
	*postgres.Store
	*Config
}

func NewServiceContext(store *postgres.Store, config *Config) *ServiceContext {
	return &ServiceContext{
		store,
		config,
	}
}
