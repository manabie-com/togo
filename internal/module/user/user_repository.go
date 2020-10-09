package user

import (
	"errors"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	GetAll() ([]User, error)
	GetUser(id uint64) (User, error)
	GetUserByEmail(email string) (User, error)
	CheckEmailExist(emai string) (bool, error)
	Create(email string, password string, maxToto int) (User, error)
}

// NewUserRepository func
func NewUserRepository(db *gorm.DB) (Repository, error) {
	return &repository{
		db: db,
	}, nil
}

type repository struct {
	db *gorm.DB
}

func (repo *repository) GetUser(id uint64) (User, error) {
	var err error
	var u User
	err = repo.db.Raw(`SELECT * FROM "users" WHERE id = ?`, id).Scan(&u).Error
	if err != nil {
		return u, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New("user not found")
	}
	return u, err
}

func (repo *repository) GetUserByEmail(email string) (User, error) {
	var err error
	var u User
	err = repo.db.Raw(`SELECT * FROM "users" WHERE email = ?`, email).Scan(&u).Error
	if err != nil {
		return u, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New("user not found")
	}
	return u, err
}

func (repo *repository) GetAll() ([]User, error) {
	var err error
	var users []User
	err = repo.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}

func (repo *repository) CheckEmailExist(email string) (bool, error) {
	var err error
	var count int64
	err = repo.db.Raw(`SELECT count(*) FROM "users" WHERE email = ?`, email).Scan(&count).Error
	if err != nil {
		return count > 0, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return count > 0, errors.New("user not found")
	}
	return count > 0, err
}
func (repo *repository) Create(email string, password string, maxToto int) (User, error) {
	user := User{Email: email, Password: password, MaxTodo: maxToto}
	repo.db.Create(&user)
	return user, nil
}
