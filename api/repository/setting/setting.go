package setting

import (
	"context"
	"database/sql"
	"time"

	"manabie/todo/models"
)

const (
	queryCreate         = `INSERT INTO setting (member_id, limit_task) VALUES ($1, $2)`
	queryUpdate         = `UPDATE setting SET limit_task = $1, updated_at = $2 WHERE id = $3`
	queryFindByMemberID = `SELECT * FROM setting WHERE member_id = $1`
	queryFindByID       = `SELECT * FROM setting WHERE id = $1`
)

type SettingRespository interface {
	Create(ctx context.Context, tx *sql.Tx, st *models.Setting) error
	Update(ctx context.Context, tx *sql.Tx, st *models.Setting) error
	FindByMemberID(ctx context.Context, tx *sql.Tx, memberID int) (*models.Setting, error)
	FindByID(ctx context.Context, tx *sql.Tx, ID int) (*models.Setting, error)
}

type settingRespository struct{}

func NewSettingRespository() SettingRespository {
	return &settingRespository{}
}

func (sr *settingRespository) Create(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	stmt, err := tx.PrepareContext(ctx, queryCreate)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, st.MemberID, st.LimitTask); err != nil {
		return err
	}

	return nil
}

func (sr *settingRespository) Update(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	stmt, err := tx.PrepareContext(ctx, queryUpdate)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, st.LimitTask, time.Now(), st.ID); err != nil {
		return err
	}

	return nil
}

func (sr *settingRespository) FindByMemberID(ctx context.Context, tx *sql.Tx, memberID int) (*models.Setting, error) {
	row := tx.QueryRowContext(ctx, queryFindByMemberID, memberID)

	st := &models.Setting{}

	if err := row.Scan(&st.ID, &st.MemberID, &st.LimitTask, &st.CreatedAt, &st.UpdateAt); err != nil {
		return nil, err
	}

	return st, nil
}

func (sr *settingRespository) FindByID(ctx context.Context, tx *sql.Tx, ID int) (*models.Setting, error) {
	row := tx.QueryRowContext(ctx, queryFindByID, ID)

	st := &models.Setting{}

	if err := row.Scan(&st.ID, &st.MemberID, &st.LimitTask, &st.CreatedAt, &st.UpdateAt); err != nil {
		return nil, err
	}

	return st, nil
}
