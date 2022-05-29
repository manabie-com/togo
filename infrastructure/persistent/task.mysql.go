package persistent

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"togo/domain/errdef"
	"togo/domain/model"
	"togo/domain/repository"
)

type taskMySQLRepository struct {
	db *sql.DB
}

func (this *taskMySQLRepository) CountTaskCreatedInDayByUser(ctx context.Context, u model.User) (int, error) {
	stmt, err := this.db.PrepareContext(ctx, "SELECT count(*) FROM todo.tbl_task WHERE created_by = ? AND date = ?")
	if err != nil {
		return 0, errdef.SystemError
	}
	currentDate := time.Now().Format("2006-01-02")
	rows, err := stmt.QueryContext(ctx, u.Id, currentDate)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var countRecord int
	for rows.Next() {

		err = rows.Scan(&countRecord)
		if err != nil {
			return 0, err
		}
	}
	return countRecord, nil
}

func (this *taskMySQLRepository) Create(ctx context.Context, u model.Task) error {
	stmt, err := this.db.PrepareContext(ctx, "CALL todo.SP_Task_Create(?, ?, ?)")
	if err != nil {
		return errdef.SystemError
	}
	rows, err := stmt.QueryContext(ctx, u.Title, u.Description, u.CreatedBy)
	if err != nil {
		return err
	}
	var errorCode int
	var msg string
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&errorCode, &msg)
		if err != nil {
			return err
		}
	}
	if errorCode != 0 {
		return errors.New(msg)
	}
	return nil
}

func NewTaskMySQLRepository(db *sql.DB) repository.TaskRepository {
	return &taskMySQLRepository{
		db: db,
	}
}
