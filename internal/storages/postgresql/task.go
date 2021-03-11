package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

// PSQLTaskRespsitory ...
type PSQLTaskRespsitory struct {
	db *sql.DB
}

type task struct {
	UserID   string
	Password string
}

// NewPSQLTaskRespsitory ...
func NewPSQLTaskRespsitory(db *sql.DB) *PSQLTaskRespsitory {
	return &PSQLTaskRespsitory{
		db: db,
	}
}

// GetTasksByUserIDAndDate ...
func (psql *PSQLTaskRespsitory) GetTasksByUserIDAndDate(ctx context.Context, userID, createdDate string) ([]*entities.Task, error) {
	var result []*entities.Task
	rowCursor, err := psql.db.Query("SELECT id, content, created_date from tasks WHERE user_id = $1 AND created_date = $2;", userID, createdDate)
	if err != nil {
		log.Println(err)
		return nil, storages.ErrInternalError
	}
	for rowCursor.Next() {
		var tmpTask entities.Task
		err := rowCursor.Scan(&tmpTask.ID, &tmpTask.Content, &tmpTask.CreatedDate)
		if err != nil {
			log.Println(err)
			return nil, storages.ErrInternalError
		}
		result = append(result, &tmpTask)
	}

	if len(result) == 0 {
		result = []*entities.Task{}
	}
	return result, nil
}

// SaveTask ...
func (psql *PSQLTaskRespsitory) SaveTask(ctx context.Context, task entities.Task) (*entities.Task, error) {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4);`
	_, err := psql.db.ExecContext(ctx, stmt, &task.ID, &task.Content, &task.UserID, &task.CreatedDate)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Cannot save to database")
	}
	// TODO: Return recent saved row to retrieve to-be-generated or default field from psql engine when saving the row,
	// for the sake of simplicity of the requirement, it's temporarily ignored here
	return &task, nil
}

// CountTasksOfUserByDate ...
func (psql *PSQLTaskRespsitory) CountTasksOfUserByDate(ctx context.Context, userID, createdDate string) (int, error) {
	var count int

	err := psql.db.QueryRow("SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2;", userID, createdDate).Scan(&count)
	if err != nil {
		return 0, storages.ErrInternalError
	}
	return count, nil
}
