// Package dbtest contains supporting code for running tests that hit the DB.
package dbtest

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"math/rand"
	"net/mail"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/data/dbmigrate"
	authenticateStore "github.com/manabie-com/togo/internal/features/authenticate/store"
	"github.com/manabie-com/togo/internal/web/auth"
	"github.com/manabie-com/togo/platform/database"
	"github.com/manabie-com/togo/platform/docker"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//go:embed sql/seed.sql
	seedDoc string
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// StartDB starts a database instance.
func StartDB() (*docker.Container, error) {
	image := "postgres:15-alpine"
	port := "5432"
	args := []string{"-e", "POSTGRES_PASSWORD=postgres"}

	c, err := docker.StartContainer(image, port, args...)
	if err != nil {
		return nil, fmt.Errorf("starting container: %w", err)
	}

	fmt.Printf("Image:       %s\n", image)
	fmt.Printf("ContainerID: %s\n", c.ID)
	fmt.Printf("Host:        %s\n", c.Host)

	return c, nil
}

// StopDB stops a running database instance.
func StopDB(c *docker.Container) {
	docker.StopContainer(c.ID)
	fmt.Println("Stopped:", c.ID)
}

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T, c *docker.Container) (*zap.SugaredLogger, *sqlx.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbM, err := database.Open(database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host,
		Name:       "postgres",
		DisableTLS: true,
	})
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Waiting for database to be ready ...")

	if err := database.StatusCheck(ctx, dbM); err != nil {
		t.Fatalf("status check database: %v", err)
	}

	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 4)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	dbName := string(b)

	if _, err := dbM.ExecContext(context.Background(), "CREATE DATABASE "+dbName); err != nil {
		t.Fatalf("creating database %s: %v", dbName, err)
	}
	dbM.Close()

	t.Log("Database ready")

	// =========================================================================

	db, err := database.Open(database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host,
		Name:       dbName,
		DisableTLS: true,
	})
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Migrate and seed database ...")

	if err := dbmigrate.Migrate(ctx, db); err != nil {
		t.Logf("Logs for %s\n%s:", c.ID, docker.DumpContainerLogs(c.ID))
		t.Fatalf("Migrating error: %s", err)
	}

	if err := dbmigrate.Seed(ctx, db); err != nil {
		t.Logf("Logs for %s\n%s:", c.ID, docker.DumpContainerLogs(c.ID))
		t.Fatalf("Seeding error: %s", err)
	}

	if err := dbmigrate.SeedCustom(ctx, db, seedDoc); err != nil {
		t.Logf("Logs for %s\n%s:", c.ID, docker.DumpContainerLogs(c.ID))
		t.Fatalf("Seeding error: %s", err)
	}

	t.Log("Ready for testing ...")

	var buf bytes.Buffer
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buf)
	log := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel),
		zap.WithCaller(true),
	).Sugar()

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()

		log.Sync()

		writer.Flush()
		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")
	}

	return log, db, teardown
}

// Test owns state for running and shutting down tests.
type Test struct {
	DB       *sqlx.DB
	Log      *zap.SugaredLogger
	Auth     *auth.Auth
	Teardown func()

	t *testing.T
}

// NewIntegration creates a database, seeds it, constructs an authenticator.
func NewIntegration(t *testing.T, c *docker.Container) *Test {
	log, db, teardown := NewUnit(t, c)

	cfg := auth.Config{
		SigningKey: "jwtsecret",
	}
	a, err := auth.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	test := Test{
		DB:       db,
		Log:      log,
		Auth:     a,
		t:        t,
		Teardown: teardown,
	}

	return &test
}

// Token generates an authenticated token for a user.
func (test *Test) Token(email string, pass string) string {
	test.t.Log("Generating token for test ...")

	addr, _ := mail.ParseAddress(email)

	store := authenticateStore.NewStore(test.Log, test.DB)
	dbUsr, err := store.GetUserByEmail(context.Background(), *addr)
	if err != nil {
		return ""
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   dbUsr.ID.String(),
			Issuer:    "todo api project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err := test.Auth.GenerateToken(claims)
	if err != nil {
		test.t.Fatal(err)
	}

	return token
}
