package postgres

import (
	"context"
	"database/sql"

	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (t *userRepository) CreateOne(ctx context.Context, exec boil.ContextExecutor, m *models.User) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (u *userRepository) GetUser(ctx context.Context, exec boil.ContextExecutor, userID int) (*models.User, error) {
	user, err := models.Users(
		qm.Load(models.UserRels.Tasks),
		models.UserWhere.ID.EQ(userID),
	).One(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.DataNotFoundError
		}
		return nil, err
	}
	return user, nil
}

func (t *userRepository) AddTask(ctx context.Context, exec boil.ContextExecutor, user *models.User, task *models.Task) error {
	return user.AddTasks(ctx, exec, true, task)
}
