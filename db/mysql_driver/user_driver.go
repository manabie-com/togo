package mysql_driver

import (
	"errors"
	"mini_project/db/model"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return users
}

func (s *database) GetUser(userID string) (*model.User, error) {
	user := model.User{ID: userID}
	err := s.db.Where("deleted = ?", 0).Find(&user).Error
	if err != nil {
		return nil, err
	}
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

func (s *database) GetListUserID() []string {
	var userID []string
	s.db.Table("user").Select("DISTINCT id").Where("deleted = ?", 0).Find(&userID)
	return userID
}

func (s *database) IsUserIDNotExists(userID string) bool {
	var user model.User
	result := s.db.Model(&model.User{}).Where("id = ?", userID).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (s *database) CreateTask(userID, taskName string) error {

	checkexist := s.IsUserIDNotExists(userID)

	if checkexist {
		return status.Error(codes.Unavailable, "userID : "+userID+" Unavailable in User List")
	}

	var tasks []model.Task
	s.db.Model(&model.Task{}).Where("user_id = ?", userID).Find(&tasks)
	task_count := 0

	if len(tasks) > 0 {
		var taskN string

		result := s.db.Table("task").Select("task_name").Where("task_name = ?", taskName).Take(&taskN)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return status.Error(codes.AlreadyExists, "taskName : "+taskName+" Already Exists in Task List")
		}

		currentTime := time.Now()

		for _, t := range tasks {
			c_y, c_m, c_d := currentTime.Date()
			t_y, t_m, t_d := t.CreatedAt.Date()

			if c_y == t_y && c_m == t_m && c_d == t_d {
				task_count++
			}
		}

		user, _ := s.GetUser(userID)

		task_count_max := user.NumberTask

		if task_count >= task_count_max {
			return status.Error(codes.OutOfRange, "number task assign to this user over task count max")
		}

	}

	task := model.Task{
		UserID:    userID,
		TaskName:  taskName,
		CreatedAt: time.Now(),
	}
	return s.db.Create(&task).Error
}
