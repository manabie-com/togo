package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/internal/must"
	"github.com/vchitai/togo/internal/server"
	"github.com/vchitai/togo/internal/service"
	"github.com/vchitai/togo/internal/store"
	"github.com/vchitai/togo/internal/utils"
	"gorm.io/gorm"
)

var (
	testDBName = "togo_test"
	ll         = l.New()
)

func TestAddToDoList(t *testing.T) {
	initDB()
	defer cleanupDB()
	var userID = "user-test"
	for _, tc := range []struct {
		name     string
		setup    func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client)
		req      map[string]interface{}
		httpCode int
		resp     map[string]interface{}
	}{
		{
			name: "success",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 1
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": userID,
				"entry": []map[string]interface{}{
					{
						"content": "a",
					},
				},
			},
			httpCode: http.StatusOK,
			resp: map[string]interface{}{
				"message": "ok",
			},
		},
		{
			name: "failed_validate_user_id",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 1
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": "",
				"entry": []map[string]interface{}{
					{
						"content": "a",
					},
				},
			},
			httpCode: http.StatusBadRequest,
			resp: map[string]interface{}{
				"code":    float64(3),
				"details": []interface{}{},
				"message": "invalid AddToDoListRequest.UserId: value length must be at least 1 runes",
			},
		},
		{
			name: "failed_validate_entry",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 1
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": userID,
				"entry":   []map[string]interface{}{},
			},
			httpCode: http.StatusBadRequest,
			resp: map[string]interface{}{
				"code":    float64(3),
				"details": []interface{}{},
				"message": "invalid AddToDoListRequest.Entry: value must contain at least 1 item(s)",
			},
		},

		{
			name: "failed_validate_entry",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 1
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": userID,
				"entry": []map[string]interface{}{
					{
						"content": "",
					},
				},
			},
			httpCode: http.StatusBadRequest,
			resp: map[string]interface{}{
				"code":    float64(3),
				"details": []interface{}{},
				"message": "invalid AddToDoListRequest.Entry[0]: embedded message failed validation | caused by: invalid ToDoEntry.Content: value length must be at least 1 runes",
			},
		},
		{
			name: "failed_limited_reached",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 0
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": userID,
				"entry": []map[string]interface{}{
					{
						"content": "a",
					},
				},
			},
			httpCode: http.StatusBadRequest,
			resp: map[string]interface{}{
				"code":    float64(3),
				"details": []interface{}{},
				"message": "you have reached daily limit",
			},
		},
		{
			name: "failed_limited_exceed",
			setup: func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client) {
				cfg.ToDoListAddLimitedPerDay = 1
				// reset count for test user
				_ = redisCli.Del(redisCli.Context(), store.BuildUserDailyUsedCount(userID, utils.RoundDate(time.Now().Add(7*time.Hour)))).Err()
			},
			req: map[string]interface{}{
				"user_id": userID,
				"entry": []map[string]interface{}{
					{
						"content": "a",
					},
					{
						"content": "a",
					},
				},
			},
			httpCode: http.StatusBadRequest,
			resp: map[string]interface{}{
				"code":    float64(3),
				"details": []interface{}{},
				"message": "you will exceed daily limit adding this list"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()
			evalRequest(ctx, t, tc.setup, tc.req, tc.httpCode, tc.resp)
		})
	}
}

func initDB() {
	cfg := configs.Load()
	db := must.ConnectMySQL(cfg.MySQL)

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

func evalRequest(ctx context.Context, t *testing.T,
	setup func(cfg *configs.Config, db *gorm.DB, redisCli *redis.Client),
	req map[string]interface{},
	httpCode int, resp map[string]interface{},
) {
	cfg := configs.Load()
	var (
		db       = must.ConnectMySQL(cfg.MySQL)
		redisCli = must.ConnectRedis(cfg.Redis)
	)
	_ = db.AutoMigrate(&models.ToDoConfig{})
	_ = db.AutoMigrate(&models.ToDo{})
	setup(cfg, db, redisCli)
	var (
		todoStore = store.NewToDo(db, redisCli)
		svc       = service.New(cfg, todoStore)
		srv       = server.NewGRPCServer().
				WithServiceServer(svc)
		gw = server.NewGatewayServer(cfg)
	)
	go func() {
		<-ctx.Done()
		srv.Stop()
	}()
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

	var buf = bytes.NewBufferString("")
	err = json.NewEncoder(buf).Encode(req)
	assert.NoError(t, err)
	testReq := httptest.NewRequest(http.MethodPost, "/to-do", buf)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, testReq)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, httpCode, res.StatusCode)
	var testResp = make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&testResp)
	assert.NoError(t, err)
	assert.Equal(t, resp, testResp)
}
