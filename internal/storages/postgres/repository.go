package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/storages"
)

type repository struct {
	Conn *gorm.DB
}

func InitRepository(conn *gorm.DB) storages.EntityRepository {
	return &repository{
		Conn: conn,
	}
}

func (repo *repository) GetTasks(userID string, createdDate string) (tasks []storages.Task, err error) {
	err = repo.Conn.Where(`user_id = ? AND created_date = ?`, userID, createdDate).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (repo *repository) InsertTask(task *storages.Task) (err error) {
	err = repo.Conn.Create(&task).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Login(userID string, pwd string) (err error) {
	err = repo.Conn.Where(`id = ? AND password = ?`, userID, pwd).First(&storages.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetUserByID(userID string) (user storages.User, err error) {
	err = repo.Conn.Where(`id = ?`, userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
