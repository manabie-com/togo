package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"togo/internal/pkg/db"
	"togo/internal/services/task/domain"
)

type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(entity *domain.User) error {
	_, err := r.DB.Conn.Model(entity).Returning("*").Insert()
	return err
}

func (r *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var entity domain.User
	err := r.DB.Conn.Model(&entity).Where("id = ?", id).First()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &entity, err
}
