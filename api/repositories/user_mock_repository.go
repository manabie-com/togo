package repositories

import "github.com/kier1021/togo/api/models"

// UserMockRepository is the mock user repository used in unit testing
type UserMockRepository struct {
	users []models.User
}

// NewUserMockRepository is the constructor for UserMockRepository
func NewUserMockRepository() *UserMockRepository {

	users := []models.User{
		{
			ID:       "620e6b6e20bdcb887326931a",
			UserName: "Test User 1",
			MaxTasks: 3,
		},
		{
			ID:       "620e6b79657f405b5f79b32d",
			UserName: "Test User 2",
			MaxTasks: 4,
		},
		{
			ID:       "620e6b7e64b5c80f08aaddcd",
			UserName: "Test User 3",
			MaxTasks: 2,
		},
	}

	return &UserMockRepository{
		users: users,
	}
}

// CreateUser returns a static ID for testing
func (repo *UserMockRepository) CreateUser(user models.User) (string, error) {
	return "620e6baff70a3fd2fc8811a0", nil
}

// GetUser return the first user that matched the filter
func (repo *UserMockRepository) GetUser(filter map[string]interface{}) (user *models.User, err error) {

	users := repo.filterUsers(filter)

	if len(users) != 0 {
		user = &users[0]
	}

	return user, nil
}

// GetUsers returns the filtered static users
func (repo *UserMockRepository) GetUsers(filter map[string]interface{}) (users []models.User, err error) {
	return repo.filterUsers(filter), nil
}

// filterUsers filter the static users based on the given filter
func (repo *UserMockRepository) filterUsers(filter map[string]interface{}) (users []models.User) {

	if len(filter) == 0 {
		return repo.users
	}

	for _, u := range repo.users {

		isEqual := false

		if id, ok := filter["_id"]; ok {
			isEqual = u.ID == id
		}

		if userName, ok := filter["user_name"]; ok {
			isEqual = u.UserName == userName
		}

		if maxTasks, ok := filter["max_tasks"]; ok {
			isEqual = u.MaxTasks == maxTasks
		}

		if isEqual {
			users = append(users, u)
		}

	}

	return users
}
