package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/usecase/interfaces"
)

type sqlRepository struct {
	db *sqlx.DB
}

var _ interfaces.UserDataSource = sqlRepository{}

func NewSQLRepository(db *sqlx.DB) interfaces.UserDataSource {
	return sqlRepository{db: db}
}

func (r sqlRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	query := "SELECT * FROM users u WHERE u.email = $1"

	if err := r.db.GetContext(ctx, u, query, email); err != nil {
		return nil, err
	}

	return u, nil
}

func (r sqlRepository) Add(ctx context.Context, u *entity.User) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO users (id, email, password, max_todo, created_at)
		VALUES (:id, :email, :password, :max_todo, :created_at)
	`

	_, err = tx.NamedExecContext(ctx, query, u)
	if err != nil {
		return err
	}

	return tx.Commit()
}
