package sqlite

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"strconv"
	"time"
	"togo-user-session-service/internal/model"
	"togo-user-session-service/internal/storage"

	"github.com/giahuyng98/togo/core-lib/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
)

type Config struct {
	Name               string
	SonyflakeStartTime time.Time
}

type SQLiteDB struct {
	db     *sql.DB
	sf     *sonyflake.Sonyflake
	config *Config
}

func NewSqliteDB(config *Config) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", config.Name)
	if err != nil {
		return nil, err
	}

	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: config.SonyflakeStartTime,
	})

	return &SQLiteDB{
		db:     db,
		config: config,
		sf:     sf,
	}, nil
}

func (s *SQLiteDB) encryptPassword(ctx context.Context, password string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(password))
	sha := sha256.New()
	sum := sha.Sum([]byte(encoded))
	return string(sum), nil
}

func (s *SQLiteDB) RegisterOrLogin(ctx context.Context, username, password string) (*model.User, error) {
	if user, err := s.Login(ctx, username, password); err == nil {
		logger.For(ctx).Info("RegisterOrLogin login success", zap.String("username", username))
		return user, nil
	}
	if user, err := s.Register(ctx, username, password); err == nil {
		logger.For(ctx).Info("RegisterOrLogin register success", zap.String("username", username))
		return user, nil
	} else {
		return nil, err
	}
}

func (s *SQLiteDB) Login(ctx context.Context, username, password string) (*model.User, error) {
	hashedPass, _ := s.encryptPassword(ctx, password)
	stm := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	rows, err := s.db.QueryContext(ctx, stm, username, hashedPass)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.UserName)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, storage.ErrWrongUserNameOrPass
}

func (s *SQLiteDB) Register(ctx context.Context, username, password string) (*model.User, error) {
	stm := "INSERT INTO users(id, username, password) VALUES(?, ?, ?)"
	hashedPass, _ := s.encryptPassword(ctx, password)

	var userID string

	if id, err := s.sf.NextID(); err != nil {
		return nil, err
	} else {
		userID = strconv.Itoa(int(id))
	}
	_, err := s.db.ExecContext(ctx, stm, userID, username, hashedPass)

	if err != nil {
		return nil, err
	}
	return &model.User{
		UserID:   userID,
		UserName: username,
	}, nil
}
