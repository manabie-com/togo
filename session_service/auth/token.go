package auth

import (
	"fmt"
	"os"
	"session_service/proto"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var redisClient *redis.Client

func InitRedis() {

	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "127.0.0.1:6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

type TokenDetails struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
}

type AccessDetails struct {
	AccessUUID string
	AccountID  uint64
}

//GenerateToken ...
func GenerateToken(accInfo *proto.AccountInfo) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 60).Unix()
	td.AccessUUID = uuid.New().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["account_id"] = accInfo.GetId()
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

//CreateAuth ...
func CreateAuth(accountID uint64, td *TokenDetails) error {

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()

	errAccess := redisClient.Set(td.AccessUUID, strconv.Itoa(int(accountID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}

//ExtractClaims ...
func ExtractClaims(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
}

//ExtractTokenMetadata ...
func ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {

	claims, err := ExtractClaims(tokenString)
	if err != nil {
		return nil, err
	}

	accessUUID, ok := claims["access_uuid"].(string)
	if !ok {
		return nil, err
	}

	accountID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["account_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &AccessDetails{
		AccessUUID: accessUUID,
		AccountID:  accountID,
	}, nil
}

//FetchAuth ...
func FetchAuth(ad *AccessDetails) (string, error) {

	accessUUID, err := redisClient.Get(ad.AccessUUID).Result()
	if err != nil {
		return "", err
	}

	return accessUUID, err
}

//DeleteAuth ...
func DeleteAuth(accessUUID string) (int64, error) {

	deleted, err := redisClient.Del(accessUUID).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
