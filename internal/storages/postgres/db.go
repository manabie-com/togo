package postgres

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/util"

	_ "github.com/lib/pq"
)

const (
	MaxLifetime  = time.Minute
	MaxIdleConns = 50
	MaxOpenConns = 50
)

type Postgres struct {
	db     *sql.DB
	logger *logs.Logger
}

func NewPostgres() *Postgres {
	logger := logs.WithPrefix("Postgres")
	p := &Postgres{
		logger: logger,
	}

	conn := util.Conf.ConnectionString()
	db, err := sql.Open(util.Conf.PostgresDriver, conn)
	if err != nil {
		logger.Panic("Connection Postgres occur error", err.Error())

	}
	p.db = db
	if err = p.pingDB(); err != nil {
		logger.Panic("Cannot ping Postgres", err.Error())
	}
	logger.Info("Connection Postgres successful", nil)

	p.db.SetConnMaxLifetime(MaxLifetime)
	p.db.SetMaxIdleConns(MaxIdleConns)
	p.db.SetMaxOpenConns(MaxOpenConns)

	defer func() {
		go func() {
			for {
				if err := db.Ping(); err != nil {
					logger.Error("Cannot ping postgres", err.Error())
				}

				time.Sleep(util.Conf.Timeout)
			}
		}()
	}()

	return p
}

func (p *Postgres) pingDB() (err error) {
	for i := 0; i < 10; i++ {
		err = p.db.Ping()

		if err == nil {
			return
		}

		p.logger.Error("Try to ping postgres", err.Error())
		time.Sleep(util.Conf.Timeout)
	}

	return err
}

// RetrieveTasks is function get list task with condition user_id, created_date and limit and offset
func (p *Postgres) RetrieveTasks(ctx context.Context, userID, createdDate string, limit, offset int) ([]*entities.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2
				LIMIT $3 OFFSET $4`
	offset = limit * (offset - 1)
	rows, err := p.db.QueryContext(ctx, stmt, userID, createdDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*entities.Task
	for rows.Next() {
		t := &entities.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask is function create task
func (p *Postgres) AddTask(ctx context.Context, t *entities.Task) error {
	err := p.execTx(ctx, func() error {
		// select user
		stmt := `SELECT max_todo FROM users WHERE id = $1`
		rows, err := p.db.QueryContext(ctx, stmt, t.UserID)
		if err != nil {
			return err
		}

		// check user and get maxtodo of user
		var maxTodo int
		var existUser bool
		for rows.Next() {
			existUser = true
			err = rows.Scan(&maxTodo)
			if err != nil {
				return err
			}
		}

		if !existUser {
			return errors.New("Don't exist User")
		}

		// count amount of task
		stmt = `SELECT count(id) FROM tasks where user_id = $1 and created_date = $2`
		rows, err = p.db.QueryContext(ctx, stmt, t.UserID, t.CreatedDate)
		var count int
		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				return err
			}
		}

		if count >= maxTodo {
			return errors.New("User's task is over limit")
		}

		// create task
		stmt = `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
		_, err = p.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// ExecTx executes a function within a database transaction
func (p *Postgres) execTx(ctx context.Context, fn func() error) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn()
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

// ValidateUser is function validate username and password of users
func (p *Postgres) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	hashPwd := hashPassword(pwd)

	row := p.db.QueryRowContext(ctx, stmt, userID, hashPwd)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {

		p.logger.Error("Invalid User", err.Error())
		return false
	}

	return true
}

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}
