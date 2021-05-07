package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const prefix = "/interview"

var Config *VecConfig
var Uptime = time.Now().String()

type VecConfig struct {
	Prefix     string
	ApiVersion string
	ServerPort string
	ServerMode string

	JwtValidMinute          int64
	AESJWTKey               string
	HMACCombinePasswordKey  string
	RedisConnectionHost     string
	RedisConnectionPassword string

	MaxLimit   int // max record on a query db.
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string

	DefaultMaxTodo uint
}

func FetchProductionEnvironmentVariables() {
	Config = NewProductionVecConfig()
	Config.GetConfig()
	Config.LoadProductionDB()
	Config.LoadProductionRedis()
}

func FetchTestEnvironmentVariables() {
	Config = NewTestVecConfig()
	Config.GetConfig()
	Config.LoadTestDB()
	Config.LoadTestRedis()
}

func (cf *VecConfig) LoadTestDB() {
	cf.DbHost = os.Getenv("TEST_DATABASE_HOST")
	cf.DbPort = os.Getenv("TEST_DATABASE_PORT")
	cf.DbUsername = os.Getenv("TEST_DATABASE_USERNAME")
	cf.DbPassword = os.Getenv("TEST_DATABASE_PASSWORD")
	cf.DbName = os.Getenv("TEST_DATABASE_NAME")
	cf.DbSSLMode = os.Getenv("TEST_DATABASE_SSL_MODE")
	cf.DbTimeZone = os.Getenv("TEST_DATABASE_TIME_ZONE")
	cf.MaxLimit, _ = strconv.Atoi(os.Getenv("TEST_DATABASE_QUERY_MAX_LIMIT"))
}

func (cf *VecConfig) LoadProductionDB() {
	cf.DbHost = os.Getenv("DATABASE_HOST")
	cf.DbPort = os.Getenv("DATABASE_PORT")
	cf.DbUsername = os.Getenv("DATABASE_USERNAME")
	cf.DbPassword = os.Getenv("DATABASE_PASSWORD")
	cf.DbName = os.Getenv("DATABASE_NAME")
	cf.DbSSLMode = os.Getenv("DATABASE_SSL_MODE")
	cf.DbTimeZone = os.Getenv("DATABASE_TIME_ZONE")
	cf.MaxLimit, _ = strconv.Atoi(os.Getenv("DATABASE_QUERY_MAX_LIMIT"))
}

func (cf *VecConfig) LoadTestRedis() {
	cf.RedisConnectionHost = fmt.Sprintf("%s:%s", os.Getenv("TEST_REDIS_HOST"), os.Getenv("TEST_REDIS_PORT"))
	cf.RedisConnectionPassword = os.Getenv("TEST_REDIS_PASSWORD")
}

func (cf *VecConfig) LoadProductionRedis() {
	cf.RedisConnectionHost = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	cf.RedisConnectionPassword = os.Getenv("REDIS_PASSWORD")
}

func NewProductionVecConfig() *VecConfig {
	cf := VecConfig{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return &cf
}

func NewTestVecConfig() *VecConfig {
	cf := VecConfig{}
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return &cf
}

func (cf *VecConfig) GetConfig() {
	port := os.Getenv("PORT")
	cf.ServerPort = port
	ginMode := strings.TrimSpace(strings.ToLower(os.Getenv("GIN_ENV")))
	switch ginMode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		gin.SetMode(ginMode)
		cf.ServerMode = ginMode
	default:
		gin.SetMode(gin.DebugMode)
		cf.ServerMode = gin.DebugMode
	}

	cf.Prefix = prefix
	cf.AESJWTKey = os.Getenv("AES_JWT_KEY")
	cf.HMACCombinePasswordKey = os.Getenv("HMAC_COMBINE_PASSWORD_KEY")
	cf.ApiVersion = os.Getenv("API_VERSION")

	cf.JwtValidMinute, _ = strconv.ParseInt(os.Getenv("JWT_VALID_MINUTE"), 10, 64)

}
