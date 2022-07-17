package task

import (
	"context"
	"database/sql"
	"time"

	"manabie/todo/models"
	"manabie/todo/pkg/apiutils"
	"manabie/todo/pkg/db"
	"manabie/todo/repository/setting"
	"manabie/todo/repository/task"

	"github.com/pkg/errors"
)

type TaskService interface {
	Index(ctx context.Context, memberID int, data string) ([]*models.Task, error)
	Show(ctx context.Context, ID int) (*models.Task, error)
	Create(ctx context.Context, memberID int, req *models.TaskCreateRequest) error
	Update(ctx context.Context, ID int, t *models.Task) error
	Delete(ctx context.Context, ID int) error
}

type service struct {
	Task    task.TaskRespository
	Setting setting.SettingRespository
}

func NewTaskService(tr task.TaskRespository, st setting.SettingRespository) TaskService {
	return &service{
		Task:    tr,
		Setting: st,
	}
}

func (s *service) Index(ctx context.Context, memberID int, date string) (tasks []*models.Task, err error) {
	if err := db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

		tasks, err = s.Task.Find(ctx, tx, memberID, date)

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *service) Show(ctx context.Context, ID int) (tk *models.Task, err error) {
	if err := db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

		tk, err = s.Task.FindByID(ctx, tx, ID, false)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if tk == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Task")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) Create(ctx context.Context, memberID int, req *models.TaskCreateRequest) error {
	target, err := time.Parse("2006-01-02", req.TargetDate)
	if err != nil {
		return errors.Wrap(apiutils.ErrInvalidValue, "target_date incorrect")
	}

	t := &models.Task{
		MemberID:   memberID,
		Content:    req.Content,
		TargetDate: target,
	}

	return db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		if _, err := tx.Exec(`LOCK TABLE task IN EXCLUSIVE MODE`); err != nil {
			return err
		}

		// Find setting by member ID
		setting, err := s.Setting.FindByMemberID(ctx, tx, t.MemberID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if setting == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Setting")
		}

		// Find for update
		tasks, err := s.Task.FindForUpdate(ctx, tx, t.MemberID, t.TargetDate.Format("2006-01-02"))
		if err != nil {
			return err
		}

		// Validate
		if setting.LimitTask <= len(tasks) {
			return errors.Wrap(apiutils.ErrInvalidValue, "user has reached the maximum daily limit")
		}

		// Create
		return s.Task.Create(ctx, tx, t)
	})
}

func (s *service) Update(ctx context.Context, ID int, t *models.Task) error {
	return db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		if _, err := tx.Exec(`LOCK TABLE task IN EXCLUSIVE MODE`); err != nil {
			return err
		}

		// Find by ID
		got, err := s.Task.FindByID(ctx, tx, ID, true)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if got == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Task")
		}

		// Set data update
		got.Content = t.Content

		// Update
		return s.Task.Update(ctx, tx, got)
	})
}

func (s *service) Delete(ctx context.Context, taskID int) error {
	return db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		// Row-Level Locks
		// Find by ID
		tk, err := s.Task.FindByID(ctx, tx, taskID, true)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if tk == nil {
			return errors.Wrap(apiutils.ErrNotFound, "Task")
		}

		// Delele
		return s.Task.Delete(ctx, tx, tk)
	})
}
