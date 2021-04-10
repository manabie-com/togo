package infra

import (
	"github.com/go-pg/pg"
	read_side "github.com/manabie-com/togo/internal/services/read-side"
	"github.com/manabie-com/togo/internal/services/tasks"
	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"
	"github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/read"
)

func ProvideUserRepo(db *pg.DB) users.UserRepo {
	return storages.NewUserRepo(db)
}

func ProvideUserConfigRepo(db *pg.DB) users.UserConfigRepo {
	return storages.NewUserConfigRepo(db)
}

func ProvideUserTaskRepo(db *pg.DB) user_tasks.UserTaskRepo {
	return storages.NewUserTaskRepo(db)
}

func ProvideReadRepo(dbSlave *DBSlave) read_side.ReadRepo {
	return read.NewReadRepo(dbSlave.DB)
}

func ProvideTaskRepo(db *pg.DB) tasks.TaskRepo {
	return storages.NewTaskRepo(db)
}
