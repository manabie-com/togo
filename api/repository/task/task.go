package task

import (
	"context"
	"database/sql"
	"time"

	"manabie/todo/models"
)

const (
	queryFindByID = `SELECT * FROM task WHERE id = $1`
	queryCreate   = `INSERT INTO task (member_id, content, target_date) VALUES ($1, $2, $3)`
	queryUpdate   = `UPDATE task SET content = $1, updated_at = $2 WHERE id = $3`
	queryDelete   = `DELETE FROM task WHERE id = $1`
)

type TaskRespository interface {
	Find(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error)
	FindForUpdate(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error)
	FindByID(ctx context.Context, tx *sql.Tx, ID int, forUpdate bool) (*models.Task, error)

	Create(ctx context.Context, tx *sql.Tx, tk *models.Task) error
	Update(ctx context.Context, tx *sql.Tx, tk *models.Task) error
	Delete(ctx context.Context, tx *sql.Tx, tk *models.Task) error
}

type taskRespository struct{}

func NewTaskRespository() TaskRespository {
	return &taskRespository{}
}

func (tr *taskRespository) Find(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error) {
	return tr.find(ctx, tx, memberID, date, false)
}

func (tr *taskRespository) FindByID(ctx context.Context, tx *sql.Tx, ID int, forUpdate bool) (*models.Task, error) {
	var query = queryFindByID
	if forUpdate {
		query = query + " FOR UPDATE"
	}

	row := tx.QueryRowContext(ctx, query, ID)

	t := &models.Task{}

	if err := row.Scan(&t.ID, &t.MemberID, &t.Content, &t.TargetDate, &t.CreatedAt, &t.UpdateAt); err != nil {
		return nil, err
	}

	return t, nil
}

func (tr *taskRespository) FindForUpdate(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error) {
	return tr.find(ctx, tx, memberID, date, true)
}

func (tr *taskRespository) Create(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	stmt, err := tx.PrepareContext(ctx, queryCreate)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, tk.MemberID, tk.Content, tk.TargetDate); err != nil {
		return err
	}

	return nil
}

func (tr *taskRespository) Update(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	stmt, err := tx.PrepareContext(ctx, queryUpdate)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, tk.Content, time.Now(), tk.ID); err != nil {
		return err
	}

	return nil
}

func (tr *taskRespository) Delete(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	stmt, err := tx.PrepareContext(ctx, queryDelete)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, tk.ID); err != nil {
		return err
	}

	return nil
}

func (tr *taskRespository) find(ctx context.Context, tx *sql.Tx, memberID int, date string, forUpdate bool) ([]*models.Task, error) {
	var (
		queryFind = `SELECT * FROM task WHERE member_id = $1`
		args      = []interface{}{memberID}
	)

	if date != "" {
		queryFind = queryFind + " AND target_date >= $2 AND target_date <= $2"
		args = append(args, date)
	}

	if forUpdate {
		queryFind = queryFind + " FOR UPDATE"
	}

	rows, err := tx.QueryContext(ctx, queryFind, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	tasks := make([]*models.Task, 0)

	for rows.Next() {

		t := &models.Task{}

		if err := rows.Scan(&t.ID, &t.MemberID, &t.Content, &t.TargetDate, &t.CreatedAt, &t.UpdateAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}
