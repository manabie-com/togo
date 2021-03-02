// +build integrate

package postgres

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_Ping_SQL(t *testing.T){
	db, err := sql.Open("pgx", "postgresql://postgres:example@localhost:5432/postgres")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err)
}

func Test_Ping_PGX(t *testing.T) {
	dbUrl := "postgresql://postgres:example@localhost:5432/postgres"
	ctx := context.Background()

	store, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal("error opening db conn", err)
	}
	defer store.Close(ctx)

	err = store.Ping(ctx)
	assert.Nil(t, err)
}