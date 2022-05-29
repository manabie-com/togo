package api

import (
	"os"
	"testing"

	"github.com/dinhquockhanh/togo/internal/pkg/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.All.DB = config.DB{
		Host:     "localhost",
		Port:     "5432",
		User:     "username",
		Password: "password",
		Name:     "togo",
		Driver:   "postgres",
	} // TODO: load config from file.

	config.All.Token.SecretKey = "Xhu091or5QiwwpxrAGxNaz09+AlVbi9t9HiJmBknsOA="
	os.Exit(m.Run())
}
