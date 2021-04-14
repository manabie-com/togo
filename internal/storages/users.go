package storages

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/go-pg/pg"
	"github.com/manabie-com/togo/internal/services/users"
)

type user struct {
	tableName struct{} `sql:"users" pg:",discard_unknown_columns"`

	ID           uuid.UUID `sql:"id,type:uuid,pk"`
	Username     string    `sql:"username,type:text"`
	Password     string    `sql:"password,type:text"`
	CreatedAt    time.Time `sql:"created_at,type:timestamp with time zone"`
	UpdatedAt    time.Time `sql:"updated_at,type:timestamp with time zone"`
	OldPasswords []string  `json:"old_passwords,type:text[],array"`
}

type userRepo struct {
	db *pg.DB
}

func NewUserRepo(
	db *pg.DB,
) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	m := &user{}
	err := r.db.WithContext(ctx).Model(m).Where("username = ?", username).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &users.User{
		ID:           m.ID,
		Username:     m.Username,
		Password:     m.Password,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		OldPasswords: m.OldPasswords,
	}, nil
}
