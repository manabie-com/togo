package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrIncorrectUsernameOrPassword = errors.New("username or password is not correct")
	ErrUserMaxTodoReached          = errors.New("user's daily-limit has been reached")
)

type Database interface {
	ValidateUser(ctx context.Context, username, password string) (*storages.User, error)
	GetTasks(ctx context.Context, usrId int, createAt time.Time) ([]*storages.Task, error)
	InsertTask(ctx context.Context, task *storages.Task) error
}

// Postgres represents a database instance for working with Postgres
type Postgres struct {
	pool *pgxpool.Pool
}

// NewPostgres create new Postgres instance
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

// init initializes pgcrypto extension, tables, indexes, ... and
// inserts default data
func (pg *Postgres) init(ctx context.Context) error {
	stmt :=
		`
		SET TIMEZONE = 'Asia/Ho_Chi_Minh';

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

		CREATE INDEX IF NOT EXISTS usr_username_pwd_hash_idx ON usr(username, pwd_hash);
		CREATE INDEX IF NOT EXISTS task_usr_id_idx ON task(usr_id);
		CREATE INDEX IF NOT EXISTS task_usr_id_create_at_idx ON task(usr_id);

		INSERT INTO usr (
			id,
			username, 
			pwd_hash, 
			max_todo
		) OVERRIDING SYSTEM VALUE VALUES (
		    1,                              
			'firstUser',
		    crypt('example', gen_salt('bf')) ,
			5
		) ON CONFLICT DO NOTHING ;

		INSERT INTO task (
			id,
			usr_id, 
			content, 
			create_at) 
		OVERRIDING SYSTEM VALUE VALUES  (
			1,
			1,
			'test 1',
			'2020-06-29'::timestamptz
		) ON CONFLICT DO NOTHING ;
		`

	_, err := pg.pool.Exec(ctx, stmt)
	if err != nil {
		return errors.Wrap(err, "Exec()")
	}
	return nil
}

// ValidateUser
func (pg *Postgres) ValidateUser(ctx context.Context, username, password string) (*storages.User, error) {
	stmt :=
		`
		SELECT 
			id,
			username,
			pwd_hash,
			max_todo
		FROM 
			usr
		WHERE 
			username = $1
			AND pwd_hash = crypt($2, pwd_hash)
		`
	row := pg.pool.QueryRow(ctx, stmt, username, password)

	usr := &storages.User{}
	err := row.Scan(&usr.Id, &usr.Username, &usr.PwdHash, &usr.MaxTodo)

	switch err {
	case nil:
		return usr, nil
	case pgx.ErrNoRows:
		return nil, ErrIncorrectUsernameOrPassword
	default:
		return nil, errors.Wrap(err, "Scan()")
	}
}

func (pg *Postgres) GetTasks(ctx context.Context, usrId int, createAt time.Time) ([]*storages.Task, error) {
	stmt :=
		`
		SELECT 
			id, usr_id, content, create_at
		FROM 
		     task
		WHERE 
		      usr_id = $1
		      AND create_at::date = $2::date
		`

	rows, err := pg.pool.Query(ctx, stmt, usrId, createAt)
	switch err {
	case nil:
		defer rows.Close()
	case pgx.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	tasks := make([]*storages.Task, 0)
	for rows.Next() {
		task := &storages.Task{}
		err := rows.Scan(
			&task.Id,
			&task.UsrId,
			&task.Content,
			&task.CreateAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Scan()")
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (pg *Postgres) InsertTask(ctx context.Context, task *storages.Task) error {
	stmt :=
		`
		INSERT INTO 
		    task (usr_id, content, create_at)
		SELECT 
		   $1, $2, $3::timestamptz
		WHERE 
			(
				SELECT count(*) FROM task
				WHERE 
					usr_id = $1
					AND create_at::date = $3::date
			) < (SELECT max_todo FROM usr WHERE id = $1)
		`

	cmd, err := pg.pool.Exec(ctx, stmt, task.UsrId, task.Content, task.CreateAt)
	if err != nil {
		return errors.Wrap(err, "Exec()")
	}

	if cmd.RowsAffected() < 1 {
		return ErrUserMaxTodoReached
	}

	return nil
}

func (pg *Postgres) Close() {
	pg.pool.Close()
}
