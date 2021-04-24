package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages/psql"
	"github.com/stretchr/testify/assert"
)

func setupStorage(t *testing.T) *psql.Storage {
	var conf psql.Config
	panicIfErr(envconfig.Process("test", &conf))
	// user=pqgotest dbname=pqgotest sslmode=verify-full
	conf.ConnString = "user=admin password=password port=5433 dbname=admin sslmode=disable"
	s, err := psql.NewStorage(conf)
	panicIfErr(err)
	return s
}

func TestPostgresql(t *testing.T) {
	s := setupStorage(t)
	err := s.AddTaskWithLimitPerDay(domain.Task{
		ID:          uuid.New().String(),
		Content:     "hello",
		UserID:      "admin",
		CreatedDate: time.Now().Format("2006-01-02"),
	}, 5)
	assert.NoError(t, err)
	fmt.Println(s.GetTasksByUserID("admin", 1, -1))
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
