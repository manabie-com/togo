package sqlite

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	_const "togo-thdung002/const"
	"togo-thdung002/entities"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) *LiteDB {
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.User{})

	return &LiteDB{DB: db}
}

func (s *LiteDB) GetTaskFilterBy(userID uint, date string) (tasks []entities.Task, err error) {
	err = s.DB.Where(entities.Task{UserID: userID, Date: date}).Find(&tasks).Error
	return
}

func (s *LiteDB) CreateTask(task *entities.Task) (uint, error) {
	err := s.DB.Create(task).Error
	if err != nil {
		log.Error(err)
		return 0, _const.ErrOnCreateTask

	}
	return task.ID, nil
}

func (s *LiteDB) CreateUser(user *entities.User) (uint, error) {
	pHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(pHashed)
	err = s.DB.Create(user).Error
	if err != nil {
		log.Error(err)
		return 0, _const.ErrOnCreateUser
	}
	return user.ID, nil
}

//func (s *LiteDB) ValidateUser(username, password string) (bool, error) {
//	var user entities.User
//
//	rs := s.DB.Where(entities.User{Username: username, Password: password}).First(&user)
//	if rs.RowsAffected > 0 {
//		return true, nil
//	}
//	return false, _const.ErrLoginFail
//}

func (s *LiteDB) GetUser(id uint) (user entities.User, err error) {
	err = s.DB.Find(&user, id).Error
	return
}
