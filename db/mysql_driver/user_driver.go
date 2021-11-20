package mysql_driver

import (
	"errors"
	"mini_project/db/model"

	"gorm.io/gorm"
)

func (s *database) AddUser(userID string, fields map[string]interface{}) (*model.User, error) {
	var user model.User
	userName, ok := fields["name"]
	if ok {
		delete(fields, "name")
	}
	s.db.FirstOrCreate(&user, map[string]interface{}{"id": userID, "name": userName})
	err := s.db.Model(&user).Updates(fields).Error
	return &user, err
}

func (s *database) UpdateUser(userID string, fields map[string]interface{}) (*model.User, error) {
	user := model.User{ID: userID}
	err := s.db.Model(&user).Updates(fields).Error
	return &user, err
}

func (s *database) CreateUser(user model.User) error {
	return s.db.Create(&user).Error
}

func (s *database) DeleteUser(user model.User) error {
	return s.db.Delete(&user).Error
}

func (s *database) IsUserNotExists(userName string) bool {
	var user model.User
	result := s.db.Model(&model.User{}).Where("name = ?", userName).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (s *database) GetUsers() []model.User {
	var users []model.User
	s.db.Where("deleted = ?", 0).Find(&users)
	var resp []model.User
	// for _, user := range users {
	// 	user.UserName = user.UserName
	// 	user.Email = user.Email
	// 	user.Phone = user.Phone
	// 	user.Whilelist.Ips = user.Whilelist.Ips
	// 	resp = append(resp, user)
	// }
	return resp
}

func (s *database) GetUser(userID string) (*model.User, error) {
	user := model.User{ID: userID}
	err := s.db.Where("deleted = ?", 0).Find(&user).Error
	if err != nil {
		return nil, err
	}
	// user.UserName = user.UserName
	// user.Phone = user.Phone
	// user.Email = user.Email
	// user.Whilelist.Ips = user.Whilelist.Ips
	return &user, err
}

func (s *database) GetUserByName(userName string) (*model.User, error) {
	var user model.User
	err := s.db.Model(&model.User{}).Where("name = ? and deleted = ?", userName, 0).Find(&user).Error
	if err != nil {
		return nil, err
	}
	// user.UserName = user.UserName
	// user.Phone = user.Phone
	// user.Email = user.Email
	// user.Whilelist.Ips = user.Whilelist.Ips
	return &user, err
}
