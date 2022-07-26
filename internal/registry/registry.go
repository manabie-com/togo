package registry

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/persistence/persister/sql"
	"github.com/trangmaiq/togo/internal/server/handler"
	"github.com/trangmaiq/togo/pkg/limiter"
)

var (
	once sync.Once
	rg   DefaultRegistry
)

type DefaultRegistry struct {
	persister *sql.Persister
	limiter   *limiter.Limiter
}

func Init(cfg *config.ToGo) (err error) {
	once.Do(func() {
		var db *sqlx.DB
		db, err = sqlx.Connect(cfg.DBDriver, cfg.DSN)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			return
		}

		rg.persister = sql.NewPersister(db)
		rg.limiter = limiter.New()
	})

	return
}

func Close() error {
	return rg.persister.Close()
}

func Registry() *DefaultRegistry {
	return &rg
}

func (r *DefaultRegistry) Persister() handler.Persister {
	return r.persister
}

func (r *DefaultRegistry) RateLimiter() handler.RateLimiter {
	return r.limiter
}
