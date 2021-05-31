package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"errors"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"log"
	"time"
)

const (
	Host = "db"
	Port = 5432
)

var (
	ErrorUserExisted = errors.New("User ID existed")
	ErrorReachLimitTask = errors.New("Reached limit task per day")
)

type DataBase struct {
	db *sql.DB
	dbName string
	user string
	password string
}

func (p *DataBase) Init(dbName, user, password string) error {
	p.dbName = dbName
	p.user = user
	p.password = password
	err := p.connect()
	go p.keepAlive()
	return err
}

func (p *DataBase) connect() error {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, p.user, p.password, p.dbName)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("Postgresql connected.")
	p.db = db
	return nil
}

func (p *DataBase) keepAlive() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			if err := p.db.Ping(); err != nil {
				log.Printf("Could not connect to Postgresql. %v\n", err)
				p.connect()
			}
		}
	}
}

func (p *DataBase) Finalize() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *DataBase) AddUser(ctx context.Context, userID, password sql.NullString) error {
	stmt := `SELECT id FROM users WHERE id = $1`
	row := p.db.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		insertStmt := `INSERT INTO "users"("id", "password") VALUES($1, crypt($2, gen_salt('bf')))`
		_, err := p.db.ExecContext(ctx, insertStmt, userID, password)
		if err != nil {
			return err
		}
	} else {
		return ErrorUserExisted
	}
	return nil
}

func (p *DataBase) ValidateUser(ctx context.Context, userID, password sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = crypt($2, password)`
	row := p.db.QueryRowContext(ctx, stmt, userID, password)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (p *DataBase) validateLimit(ctx context.Context, t *storages.Task, limit int) error {
	stmt := `SELECT COUNT(id) FROM tasks WHERE user_id = $1 AND target_date = $2`
	row := p.db.QueryRowContext(ctx, stmt, t.UserID, t.TargetDate)
	taskNum := 0
	err := row.Scan(&taskNum)
	if err == nil && taskNum >= limit {
		return ErrorReachLimitTask
	}
	return nil
}

func (p *DataBase) AddTask(ctx context.Context, t *storages.Task, limit int) error {
	if err := p.validateLimit(ctx, t, limit); err != nil {
		return err
	}

	stmt := `INSERT INTO tasks (id, content, user_id, status, created_date, target_date) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := p.db.ExecContext(ctx, stmt, t.ID, t.Content, t.UserID, t.Status, t.CreatedDate, t.TargetDate)
	if err != nil {
		return err
	}

	return nil
}

func (p *DataBase) UpdateTask(ctx context.Context, t *storages.Task, limit int) error {
	if err := p.validateLimit(ctx, t, limit); err != nil {
		return err
	}

	stmt := `UPDATE tasks SET content = $1, status = $2, target_date = $3 WHERE id = $4 AND user_id = $5`
	_, err := p.db.ExecContext(ctx, stmt, t.Content, t.Status, t.TargetDate, t.ID, t.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (p *DataBase) RetrieveTasks(ctx context.Context, userID sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1`
	rows, err := p.db.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate, &t.TargetDate)
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
