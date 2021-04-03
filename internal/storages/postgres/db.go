package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/util"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewPostgres() *Postgres {
	logger := logs.WithPrefix("Postgres")
	conn := GetConnString(*util.Conf)
	// conn := os.Getenv("CONNECTIONSTRING")
	logger.Info(conn)
	db, err := sql.Open(util.Conf.PostgresDriver, conn)
	if err != nil {
		logger.Error("Connection Postgres occur error", zap.Any("error", err.Error()))
		return nil
	}

	if err = db.Ping(); err != nil {
		logger.Error("Cannot ping postgres", zap.Any("error", err.Error()))
	}

	logger.Info("Connection Postgres successful")

	defer func() {
		go func() {
			for {
				if err := db.Ping(); err != nil {
					logger.Error("Cannot ping postgres", zap.Any("error", err.Error()))
				}

				time.Sleep(time.Second * 10)
			}

		}()
	}()

	return &Postgres{
		db:     db,
		logger: logger,
	}
}

func GetConnString(c util.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"db", "5432", "root", "secret", "todo_app", "disable")
}

func (p *Postgres) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error) {
	return nil, nil
}

func (p *Postgres) AddTask(ctx context.Context, t *entities.Task) error {
	return nil
}

func (p *Postgres) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	p.logger.Info(userID.String)
	p.logger.Info(pwd.String)
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.db.QueryRowContext(ctx, stmt, userID.String, pwd.String)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {

		p.logger.Info(err.Error())
		return false
	}

	return true
}
