package test

import (
	"ansidev.xyz/pkg/db"
	gormPkg "ansidev.xyz/pkg/gorm"
	"ansidev.xyz/pkg/log"
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	log.InitLogger("console")
}

type PostgresTestUtil struct {
}

type PostgresTestSuite struct {
	suite.Suite
	SqlDb        *sql.DB
	GormDb       *gorm.DB
	migrationDir string
}

func (s *PostgresTestSuite) initDbConnection(dbConfig db.SqlDbConfig) {
	s.SqlDb = db.NewPostgresClient(dbConfig)
	dialector := postgres.New(postgres.Config{
		Conn:                 s.SqlDb,
		PreferSimpleProtocol: true,
	})
	s.GormDb = gormPkg.InitGormDb(dialector)
}

func (s *PostgresTestSuite) SetupSuite() {
	s.migrationDir = "../../migration"
	dbConfig, err := GetTestDbConfig()
	require.NoError(s.T(), err)

	s.initDbConnection(dbConfig)
}

func (s *PostgresTestSuite) BeforeTest(suite string, method string) {
	log.Info(fmt.Sprintf("Suite: %s, Before running %s", suite, method))
	err := goose.Up(s.SqlDb, s.migrationDir)
	require.NoError(s.T(), err)
}

func (s *PostgresTestSuite) AfterTest(suite string, method string) {
	log.Info(fmt.Sprintf("Suite: %s, After running %s", suite, method))
	err := goose.Reset(s.SqlDb, s.migrationDir)
	require.NoError(s.T(), err)
}

func (s *PostgresTestSuite) TearDownSuite() {
	err := s.SqlDb.Close()
	require.NoError(s.T(), err)
}

func (s *PostgresTestSuite) SetMigrationDir(migrationDir string) {
	s.migrationDir = migrationDir
}
