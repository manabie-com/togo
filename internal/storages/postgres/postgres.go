package postgres

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

// PDB for working with postgres
type PDB struct {
	DB *pgxpool.Pool
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (pdb PDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]entities.Task, error) {
	var tasks []entities.Task
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	err := pgxscan.Select(ctx, pdb.DB, &tasks, stmt, userID, createdDate)
	return tasks, err
}

// AddTask adds a new task to DB
func (pdb PDB) AddTask(ctx context.Context, t entities.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := pdb.DB.Exec(ctx, stmt, t.ID, t.Content, t.UserID, t.CreatedDate)
	return err // if err == nil -> return nil
}

// ValidateUser returns boolean if match userID AND password
func (pdb PDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	hashedPass, err := pdb.GetHashedPass(ctx, userID)
	if err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pwd.String)); err != nil {
		return false
	}
	return true
}

// GetHashedPass returns hashed password if match userID AND password
func (pdb PDB) GetHashedPass(ctx context.Context, userID sql.NullString) (string, error) {
	var hashedPass string
	stmt := `SELECT password FROM "users" WHERE id = $1`
	row := pdb.DB.QueryRow(ctx, stmt, userID)
	err := row.Scan(&hashedPass)
	if err != nil {
		return hashedPass, err
	}
	return hashedPass, nil
}

// GetMaxTaskTodo get the number of limit task accordinate with userID
func (pdb PDB) GetMaxTaskTodo(ctx context.Context, userID string) (int, error) {
	var maxTask int
	stmt := `SELECT max_todo FROM "users" WHERE id = $1`
	row := pdb.DB.QueryRow(ctx, stmt, userID)
	err := row.Scan(&maxTask)
	if err != nil {
		return maxTask, err
	}
	return maxTask, nil
}
