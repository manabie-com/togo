package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/internal/must"
	"github.com/vchitai/togo/internal/server"
	"github.com/vchitai/togo/internal/service"
	"github.com/vchitai/togo/internal/store"
)

var (
	testDBName = "togo_test"
	ll         = l.New()
)

func initDB() {
	cfg := configs.Load()
	db := must.ConnectMySQL(cfg.MySQL)
	_ = must.ConnectRedis(cfg.Redis)

	// Delete the database if exists
	if err := db.Exec("DROP DATABASE IF EXISTS " + testDBName).Error; err != nil {
		panic(err)
	}

	if err := db.Exec("CREATE DATABASE " + testDBName).Error; err != nil {
		panic(err)
	}

	// So that the server will use the right db
	_ = os.Setenv("MYSQL__DATABASE", testDBName)
}

func cleanupDB() {
	cfg := configs.Load()
	db := must.ConnectMySQL(cfg.MySQL)

	// Delete the database if exists
	if err := db.Exec("DROP DATABASE IF EXISTS " + testDBName).Error; err != nil {
		panic(err)
	}
}

func TestAddAndGetDiagram(t *testing.T) {
	initDB()
	defer cleanupDB()
	cfg := configs.Load()
	var (
		db       = must.ConnectMySQL(cfg.MySQL)
		redisCli = must.ConnectRedis(cfg.Redis)
	)
	_ = db.AutoMigrate(&models.ToDoConfig{})
	_ = db.AutoMigrate(&models.ToDo{})
	var (
		todoStore = store.NewToDo(db, redisCli)
		svc       = service.New(cfg, todoStore)
		srv       = server.NewGRPCServer().
				WithServiceServer(svc)
		gw = server.NewGatewayServer(cfg)
	)
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
		if err != nil {
			ll.Panic("error listening to address", l.Int("address", cfg.GRPCAddress), l.Error(err))
			return
		}
		ll.Info("GRPC server start listening", l.Int("GRPC address", cfg.GRPCAddress))
		_ = srv.Serve(listener)
	}()

	mux, err := gw.GetGRPCMux(svc)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/to-do", bytes.NewBufferString(`{
	"user_id": "user-id",
    "entry": [
        {
            "content": "a"
        }
    ]
}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	if string(data) != `{"message":"ok"}` {
		t.Errorf("expected ok got %v", string(data))
	}
}
