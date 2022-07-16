package setting

import (
	"context"
	"database/sql"

	"manabie/todo/models"
	"manabie/todo/pkg/apiutils"
	"manabie/todo/pkg/db"
	"manabie/todo/repository/setting"

	"github.com/pkg/errors"
)

type SettingService interface {
	Show(ctx context.Context, memberID int) (*models.Setting, error)

	Create(ctx context.Context, memberID int, req *models.SettingCreateRequest) error
	Update(ctx context.Context, settingID int, req *models.SettingUpdateRequest) error
}

type service struct {
	Setting setting.SettingRespository
}

func NewSettingService(sr setting.SettingRespository) SettingService {
	return &service{
		Setting: sr,
	}
}

func (s *service) Show(ctx context.Context, memberID int) (st *models.Setting, err error) {
	if err := db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

		st, err = s.Setting.FindByMemberID(ctx, tx, memberID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if st == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Setting")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return st, nil
}

func (s *service) Create(ctx context.Context, memberID int, req *models.SettingCreateRequest) error {
	return db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		// Check Setting Exit
		ext, err := s.Setting.FindByMemberID(ctx, tx, memberID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if ext != nil {
			return errors.Wrap(apiutils.ErrInvalidValue, "Setting exists")
		}

		st := &models.Setting{
			MemberID:  memberID,
			LimitTask: req.LimitTask,
		}

		return s.Setting.Create(ctx, tx, st)
	})
}

func (s *service) Update(ctx context.Context, settingID int, req *models.SettingUpdateRequest) error {
	return db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		st, err := s.Setting.FindByID(ctx, tx, settingID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if st == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Setting")
		}

		st.LimitTask = req.LimitTask

		return s.Setting.Update(ctx, tx, st)
	})
}
