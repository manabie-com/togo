package entity

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"manabie-com/togo/global"
	pkg_logrus "manabie-com/togo/pkg/logger"
	"manabie-com/togo/util"
	"net/url"
	"sync"
)

func InitializeDb() {
	SetDbProvider(&Gorm{
		DbHost:     global.Config.DbHost,
		DbPort:     global.Config.DbPort,
		DbUsername: global.Config.DbUsername,
		DbPassword: global.Config.DbPassword,
		DbName:     global.Config.DbName,
		DbSSLMode:  global.Config.DbSSLMode,
		DbTimeZone: global.Config.DbTimeZone,
		LogLevel:   logger.Info,
		once:       sync.Once{},
		db:         nil,
	})
	fmt.Println()
}

var dbProvider DbProvider

type DbProvider interface {
	Db() *gorm.DB
}

// SetDbProvider sets the provider to get a gorm db connection.
func SetDbProvider(provider DbProvider) {
	dbProvider = provider
}

// HasDbProvider returns true if a db provider exists.
func HasDbProvider() bool {
	return dbProvider != nil
}

// Db returns a database connection.
func Db() *gorm.DB {
	return dbProvider.Db()
}

// UnscopedDb returns an unscoped database connection.
func UnscopedDb() *gorm.DB {
	return Db().Unscoped()
}

type Gorm struct {
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string
	LogLevel   logger.LogLevel // "gorm.io/gorm/logger"

	once sync.Once
	db   *gorm.DB
}

func (g *Gorm) GetDns() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		url.QueryEscape(util.TrimSpaceToLower(g.DbHost)),
		url.QueryEscape(util.TrimSpaceToLower(g.DbPort)),
		url.QueryEscape(util.TrimSpaceToLower(g.DbUsername)),
		url.QueryEscape(util.TrimSpaceToLower(g.DbName)),
		g.DbPassword,
		url.QueryEscape(util.TrimSpaceToLower(g.DbSSLMode)),
		url.QueryEscape(util.TrimSpaceToLower(g.DbTimeZone)),
	)
}

// Db returns the gorm db connection.
func (g *Gorm) Db() *gorm.DB {
	g.once.Do(g.Connect)

	if g.db == nil {
		log.Fatal("entity: database not connected")
	}

	return g.db
}

// Connect creates a new gorm db connection.
func (g *Gorm) Connect() {
	db, err := gorm.Open(postgres.Open(g.GetDns()), &gorm.Config{
		Logger: logger.Default.LogMode(g.LogLevel),
	})

	if err != nil || db == nil {
		if err != nil || db == nil {
			fmt.Println("Db connect:", err)
			pkg_logrus.Lgrus.Fatal("Db connect: %v", err)
		}
	} else {
		dbInstance, _ := db.DB()
		dbInstance.SetMaxIdleConns(5)
		dbInstance.SetMaxOpenConns(10)
		color.Green("Yay! " + g.DbName + " Database Connected!")
		color.Green("Database Host: " + g.DbHost)
	}

	g.db = db
}

// Close closes the gorm db connection.
func (g *Gorm) Close() {
	if g.db != nil {
		dbInstance, _ := g.db.DB()
		if err := dbInstance.Close(); err != nil {
			log.Fatal(err)
		}

		g.db = nil
	}
}
