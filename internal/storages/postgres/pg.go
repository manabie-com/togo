package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	Host string
	Port string
	Usr  string
	Pwd  string
	Db   string
}

func (c *Config) toConnStr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.Usr, c.Pwd, c.Host, c.Port, c.Db)
}

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context) (*Postgres, error) {
	var connStr string
	switch v := ctx.Value("config").(type) {
	case *Config:
		connStr = v.toConnStr()
	default:
		return nil, errors.New("no config")
	}

	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "Connect()")
	}

	pg := &Postgres{
		pool: pool,
	}

	if err := pg.init(ctx); err != nil {
		return nil, errors.Wrap(err, "init()")
	}

	return pg, nil
}

func (pg *Postgres) init(ctx context.Context) error {
	stmt :=
		`
		CREATE EXTENSION IF NOT EXISTS pgcrypto;

		CREATE TABLE IF NOT EXISTS usr (
		    id 			int GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
		    username	varchar(36) NOT NULL UNIQUE ,
		    pwd_hash 	text NOT NULL ,
		    max_todo 	int NOT NULL DEFAULT 5 CHECK ( max_todo >= 0 )
		);
		CREATE TABLE IF NOT EXISTS task (
		  	id 			int GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
		  	usr_id 		int NOT NULL REFERENCES usr(id),
		  	content 	text NOT NULL ,
		  	create_at	timestamptz NOT NULL
		);

		CREATE INDEX IF NOT EXISTS usr_username_password_id ON usr(username, pwd_hash);

		INSERT INTO usr (
			username, 
			pwd_hash, 
			max_todo
		) VALUES (
			'firstUser',
		    crypt('example', gen_salt('bf')) ,
			5
		);
		`

	_, err := pg.pool.Exec(ctx, stmt)
	if err != nil {
		return errors.Wrap(err, "Exec()")
	}
	return nil
}

func (pg *Postgres) validateUser(ctx context.Context, username, password string) error {
	stmt :=
		`
		SELECT exists (
		    SELECT 
				*
			FROM 
				usr
			WHERE 
				username = $1
				AND pwd_hash = crypt($2, pwd_hash)
		)
		`
	row := pg.pool.QueryRow(ctx, stmt, username, password)

	var valid bool
	if err := row.Scan(valid); err != nil {
		return errors.Wrap(err, "Scan()")
	}

	if valid {
		return nil
	} else {
		return errors.New("username or password is not correct")
	}
}

func (pg *Postgres) Close() {
	pg.pool.Close()
}