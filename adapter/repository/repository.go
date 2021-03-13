package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/valonekowd/togo/adapter/repository/task"
	"github.com/valonekowd/togo/adapter/repository/user"
	"github.com/valonekowd/togo/usecase/interfaces"
)

type Config struct {
	db *sqlx.DB
	// mongoDB *mongo.Database
	// es      *elasticsearch.Client
	// rdb     *redis.Client
}

type ConfigOption func(*Config)

func WithDB(db *sqlx.DB) ConfigOption {
	return func(c *Config) {
		c.db = db
	}
}

func New(options ...ConfigOption) interfaces.DataSource {
	c := &Config{}

	for _, o := range options {
		o(c)
	}

	var taskRepository interfaces.TaskDataSource
	{
		taskRepository = task.NewSQLRepository(c.db)
	}

	var userRepository interfaces.UserDataSource
	{
		userRepository = user.NewSQLRepository(c.db)
	}

	return interfaces.DataSource{
		Task: taskRepository,
		User: userRepository,
	}
}
