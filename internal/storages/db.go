package storages

import (
	"github.com/google/martian/log"
	"gorm.io/gorm"
)

// LiteDB for working with sqllite
type Store struct {
	*gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	db = db.Debug()
	s := &Store{DB: db}
	m := []interface{}{
		&User{},
		&Task{},
	}
	if err := s.AutoMigrate(m...); err != nil {
		panic(err)
	}
	return s
}

func (s *Store) AddUser(userID, password string, maxTodo int32) error {
	return s.Model(&User{}).Create(&User{ID: userID, Password: password, MaxTodo: maxTodo}).Error
}

func (s *Store) CountTasks(userID, date string) (int32, error) {
	var numOfTask int32
	stmt := `SELECT COUNT(t.id) FROM tasks t WHERE t.id = ? AND t.created_date = ?`
	err := s.DB.Raw(stmt, userID, date).Scan(&numOfTask).Error
	if err != nil {
		return -1, err
	}
	return numOfTask, nil
}

func (s *Store) GetMaxTodo(userID string) (int32, error) {
	user := &User{}
	err := s.Model(user).Select("max_todo").Where("id = ?", userID).First(user).Error
	if err != nil {
		return -1, err
	}
	return user.MaxTodo, err
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (s *Store) RetrieveTasks(userID, createdDate string) ([]*Task, error) {
	tasks := make([]*Task, 0)
	err := s.Where(&Task{UserID: userID, CreatedDate: createdDate}).Find(&tasks).Error
	return tasks, err
}

// AddTask adds a new task to DB
func (s *Store) AddTask(t *Task, callback func(string, string) error) error {
	return s.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(t).Error
		if err != nil {
			return err
		}
		err = callback(t.UserID, t.CreatedDate)
		if err != nil {
			return err
		}
		return nil
	})
}

// ValidateUser returns tasks if match userID AND password
func (s *Store) ValidateUser(userID, pwd string) bool {
	user := &User{
		ID: userID,
		Password: pwd,
	}
	err := s.DB.Model(&User{}).Select("id").Where(user).First(user).Error
	if err != nil {
		log.Errorf("error while getting user from id and password - %s", err.Error())
		return false
	}
	return true
}