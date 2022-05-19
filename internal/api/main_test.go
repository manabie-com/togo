package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	db "task-manage/internal/db/sqlc"
	"task-manage/internal/utils"
	"testing"
	"time"
)

func newTestServer(t *testing.T) *Server {
	config, err := utils.LoadConfig("../../")
	config.TokenSymmetricKey = utils.RandomString(32)
	config.AccessTokenDuration = time.Minute
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := db.New(conn)
	server, err := NewServer(config, queries)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
