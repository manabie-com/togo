package redis

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
	"togo-user-session-service/internal/model"
	"togo-user-session-service/internal/storage"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr      string
	Password  string
	DB        int
	TokenTTL  time.Duration
	SecretKey string
}

type RedisStorage struct {
	client *redis.Client
	config *Config
}

func generateTokenKey(userID string) string {
	return fmt.Sprintf("togo:user:%s:token", userID)
}

func NewRedisStorage(config *Config) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: client,
		config: config,
	}, nil
}

// https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
func (r *RedisStorage) encrypt(user *model.User) (string, error) {
	c, err := aes.NewCipher([]byte(r.config.SecretKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	raw, _ := json.Marshal(user)

	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, raw, nil)), nil
}

// https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
func (r *RedisStorage) decrypt(token string) (*model.User, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("base61.Decode %v", err)
	}

	c, err := aes.NewCipher([]byte(r.config.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("aes.NewChipher %v", err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("invalid ciphertext")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("gcm.Open %v", err)
	}

	var user model.User
	err = json.Unmarshal(plaintext, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisStorage) CreateToken(ctx context.Context, user *model.User) (string, error) {

	encrypted, err := r.encrypt(user)
	if err != nil {
		return "", err
	}

	if err := r.client.Set(ctx, generateTokenKey(user.UserID), encrypted, r.config.TokenTTL).Err(); err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func (r *RedisStorage) VerifyToken(ctx context.Context, token string) (*model.User, error) {
	userDecrypted, err := r.decrypt(token)
	if err != nil {
		return nil, err
	}

	val, err := r.client.Get(ctx, generateTokenKey(userDecrypted.UserID)).Result()

	if err != nil {
		return nil, err
	}

	if val == token {
		return userDecrypted, nil
	}
	return nil, storage.ErrTokenInvalid
}
