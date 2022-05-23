package initializer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/manabie-com/togo/migration"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GlobalParams struct contains all configurations dependencies
type GlobalParams struct {
	Enviroments *EnviromentParams
	Components  *ComponentParams
}

// EnviromentParams structure contains .env information
type EnviromentParams struct {
	DbHost     string `envconfig:"DB_HOST"`
	DbName     string `envconfig:"DB_NAME"`
	DbPort     string `envconfig:"DB_PORT"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
}

// ComponentParams structure contains used components parameters
type ComponentParams struct {
	Db        *gorm.DB
	GinEngine *gin.Engine
}

func newGlobalParams() GlobalParams {
	return GlobalParams{
		Enviroments: &EnviromentParams{},
		Components:  &ComponentParams{},
	}
}

func (gl *GlobalParams) LoadAllConfig(envFileName string) (err error) {
	// Dependencies ordered call
	if err = gl.LoadEnviromemts(envFileName); err != nil {
		return
	}
	if err = gl.LoadComponents(gl.Enviroments); err != nil {
		return
	}
	return nil
}

func (gl *GlobalParams) LoadEnviromemts(envFileName string) (err error) {
	// Load .env file parameters
	wd, _ := os.Getwd()
	if err = godotenv.Load(filepath.Join(wd, envFileName)); err != nil {
		return fmt.Errorf("Load enviroments file %s failed", envFileName)
	}

	// Processing Enviroment parameters
	err = envconfig.Process("", gl.Enviroments)
	if err != nil {
		return fmt.Errorf("Processing enviroment parameters failed")
	}
	return nil
}

func (gl *GlobalParams) LoadComponents(env *EnviromentParams) (err error) {
	var db *gorm.DB
	var ginRouter *gin.Engine

	// Loading Database connection informations
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName)
	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		return fmt.Errorf("Error when open Database: err=%s", err)
	}
	gl.Components.Db = db
	err = migration.Migrate(gl.Components.Db)
	if err != nil {
		return
	}

	// Loading Gin configurations
	ginRouter = gin.Default()
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(gin.Logger())
	ginRouter.Use(CORSMiddleware())
	gl.Components.GinEngine = ginRouter

	return nil
}

func (gl *GlobalParams) GinRouter() *gin.Engine {
	return gl.Components.GinEngine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept, Origin, Cache-Control, X-Requested-With, User-Agent, Accept-Language, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
