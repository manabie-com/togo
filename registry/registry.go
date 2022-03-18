package registry

import (
	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/pkg/database"
)

type Registry struct {
	Config config.Config
	DB     *database.Database
}

// New ...
func New(c config.Config) (*Registry, error) {
	return &Registry{
		Config: c,
		DB:     database.New(c.Databases),
	}, nil
}
