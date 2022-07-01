package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lawtrann/togo"
	_ "github.com/lib/pq"
)

type TodoDB struct {
	DB *sql.DB
}

func GetDns() string {
	host := os.Getenv("POSTGRESQL_HOST")
	port, err := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	if err != nil {
		log.Fatalf("Postgres port %s is not valid", os.Getenv("POSTGRESQL_PORT"))
		os.Exit(1)
	}
	user := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DATABASE")

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func NewTodoDB() (togo.TodoDB, error) {
	db, err := sql.Open("postgres", GetDns())
	if err != nil {
		panic(err)
	}

	return &TodoDB{
		DB: db,
	}, nil
}

// Get user by username
func (td *TodoDB) GetUserByName(userName string) (*togo.User, error) {
	var result togo.User

	// Query for a value based on a single row.
	err := td.DB.QueryRow("SELECT id, user_name, limited_per_day FROM users WHERE user_name = $1",
		userName).Scan(&result.ID, &result.UserName, &result.LimitedPerDay)

	if err != nil && err != sql.ErrNoRows {
		return &togo.User{}, err
	}

	return &result, nil
}

// Check if exceed a limited per day
func (td *TodoDB) IsExceedPerDay(u togo.User) (bool, error) {
	var result bool

	// Query for a value based on a single row.
	err := td.DB.QueryRow("SELECT (COUNT(*) >= $1) is_exceed FROM todos WHERE user_id = $2 and create_date >= NOW()::date",
		u.LimitedPerDay, u.ID).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return result, nil
}

// Adding new todo task to user if not exceed a limited per day
func (td *TodoDB) AddTodoByUser(u *togo.User, t *togo.Todo, uFlag bool) error {

	// Begin
	ctx := context.Background()
	tx, err := td.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create user if not exist
	if !uFlag {
		stmt := `INSERT INTO users(
					user_name,
					limited_per_day,
					create_date,
					create_user)
		  		  VALUES(
					$1,
					$2,
					NOW(),
					'admin')
				  RETURNING id`

		// res, err := tx.ExecContext(ctx, stmt, &u.UserName, limited)
		err = tx.QueryRowContext(ctx, stmt, &u.UserName, &u.LimitedPerDay).Scan(&t.UserID)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
	}

	stmt := `INSERT INTO todos(
				todo_id,
				user_id,
				description,
				create_date,
				create_user)
			VALUES(
				(SELECT (COUNT(*) + 1) FROM todos WHERE user_id = $1),
				$1,
				$2,
				NOW(),
				'admin')
			RETURNING todo_id`

	err = tx.QueryRowContext(ctx, stmt, t.UserID, t.Description).Scan(&t.TodoID)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return err
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
