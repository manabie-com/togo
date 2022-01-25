package infrastructure

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/database/ent/migrate"
	"time"
)

type DB struct {
	*ent.Client
}

const (
	mysqlConnStrFmt = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&loc=%s"
)

func NewDB(cfg *configs.Config) (*DB, func(), error) {
	client, err := ent.Open(dialect.MySQL, fmt.Sprintf(mysqlConnStrFmt, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "Asia%2fBangkok&parseTime=true"))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := client.Close(); err != nil {
			logrus.Error(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	if err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
		migrate.WithFixture(true),
	); err != nil {
		logrus.Panicf("failed creating schema resources: %v", err)
	}
	cancel()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, cleanup, err
	}
	return &DB{client}, cleanup, nil
}
