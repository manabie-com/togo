package database

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func NewDockerConnector(scriptPaths []string) Connector {
	return &dockerConnector{
		scriptPaths: scriptPaths,
	}
}

type dockerConnector struct {
	db          *sql.DB
	resource    *dockertest.Resource
	scriptPaths []string
}

func (p *dockerConnector) Open(cfg *Config) error {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return err
	}

	// pulls an image, creates a container based on it and runs it
	// dbPort := strings.Split(cfg.Address, ":")[1]
	opts := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", cfg.Password),
		},
		// ExposedPorts: []string{dbPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			// docker.Port(dbPort): {
			// 	{
			// 		HostIP:   "0.0.0.0",
			// 		HostPort: dbPort,
			// 	},
			// },
		},
	}

	p.resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		return err
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	return pool.Retry(func() error {
		for _, scriptPath := range p.scriptPaths {
			err := p.runSqlScript(p.resource.Container.ID, cfg.UserName, cfg.Password, scriptPath)
			if err != nil {
				return err
			}
		}
		address := fmt.Sprintf(
			"%s:%s",
			"127.0.0.1",
			p.resource.GetPort("3306/tcp"),
		)
		return p.connectDB(cfg, address)
	})
}

func (p *dockerConnector) Close() error {
	p.db.Close()
	return p.resource.Close()
}

func (p *dockerConnector) GetDB() *sql.DB {
	return p.db
}

func (p *dockerConnector) connectDB(cfg *Config, address string) error {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.UserName,
		cfg.Password,
		address,
		cfg.Database,
	)
	db, err := sql.Open("mysql", str)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(cfg.NumberIdleConns)
	db.SetMaxOpenConns(cfg.NumberMaxConns)
	db.SetConnMaxLifetime(0)
	p.db = db
	return nil
}

func (p *dockerConnector) runSqlScript(containerId, username, password, scriptPath string) error {
	cmdStr := fmt.Sprintf(
		"docker exec -i %s mysql -u %s -p%s < %s", containerId, username, password, scriptPath,
	)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
