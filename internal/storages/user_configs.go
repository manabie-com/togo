package storages

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/go-pg/pg"
	"github.com/manabie-com/togo/internal/services/users"
)

type userConfig struct {
	tableName struct{} `sql:"user_configs" pg:",discard_unknown_columns"`

	ID        int64     `sql:"id,type:bigint,pk"`
	UserID    uuid.UUID `sql:"user_id,type:uuid,fk,unique"`
	TaskLimit int       `sql:"task_limit,type:int"`
	IsActive  bool      `sql:"is_active,type:boolean"`
	CreatedAt time.Time `sql:"created_at,type:timestamp with time zone"`
	UpdatedAt time.Time `sql:"updated_at,type:timestamp with time zone"`
}

type userConfigRepo struct {
	db *pg.DB
}

func NewUserConfigRepo(
	db *pg.DB,
) *userConfigRepo {
	return &userConfigRepo{
		db: db,
	}
}

func (r *userConfigRepo) GetByUserID(ctx context.Context, userID uuid.UUID) (*users.UserConfig, error) {
	m := &userConfig{}
	err := r.db.WithContext(ctx).Model(m).Where("user_id = ? and is_active is true", userID).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &users.UserConfig{
		ID:        m.ID,
		UserID:    m.UserID,
		TaskLimit: m.TaskLimit,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (r *userConfigRepo) UpdateConfigByUserID(ctx context.Context, usercfg *users.UserConfig) error {
	if usercfg == nil {
		return nil
	}

	return r.db.WithContext(ctx).Update(&userConfig{
		ID:        usercfg.ID,
		UserID:    usercfg.UserID,
		TaskLimit: usercfg.TaskLimit,
		IsActive:  usercfg.IsActive,
		CreatedAt: usercfg.CreatedAt,
		UpdatedAt: usercfg.UpdatedAt,
	})
}
