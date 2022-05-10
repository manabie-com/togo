package persistence

import (
	"testing"

	"github.com/jfzam/togo/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.UserName = "user1"
	user.Password = "password0001"
	user.TaskLimitPerDay = 1

	repo := NewUserRepository(conn)

	u, saveErr := repo.SaveUser(&user)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.UserName, "user1")
	assert.EqualValues(t, u.TaskLimitPerDay, 1)
	//The pasword is supposed to be hashed, so, it should not the same the one we passed:
	assert.NotEqual(t, u.Password, "password0001")
}

// Failure can be due to duplicate username, etc
// Here, we will attempt saving a user that is already saved
func TestSaveUser_Failure(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.UserName = "testuser1"
	user.Password = "testpassword1"
	user.TaskLimitPerDay = 1

	repo := NewUserRepository(conn)
	u, saveErr := repo.SaveUser(&user)
	dbMsg := map[string]string{
		"username_taken": "username already taken",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	user, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUser(user.ID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.UserName, user.UserName)
	assert.EqualValues(t, u.TaskLimitPerDay, user.TaskLimitPerDay)
}

func TestGetUsers_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the users
	_, err = seedUsers(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	users, getErr := repo.GetUsers()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByUsernameAndPassword_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = &entity.User{
		UserName: "testuser1",
		Password: "testpassword1",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByUsernameAndPassword(user)

	assert.Nil(t, getErr)
	assert.EqualValues(t, user.UserName, u.UserName)
	//Note, the user password from the database should not be equal to a plane password, because that one is hashed
	assert.NotEqual(t, user.Password, u.Password)
}

func TestGetUserByUsernameAndPasswordWrongUsername_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = &entity.User{
		UserName: "testuserX",
		Password: "testpassword1",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByUsernameAndPassword(user)

	dbMsg := map[string]string{
		"no_user": "user not found",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, getErr)
}

func TestGetUserByUsernameAndPasswordWrongPassword_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = &entity.User{
		UserName: "testuser1",
		Password: "testpasswordX",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByUsernameAndPassword(user)

	dbMsg := map[string]string{
		"incorrect_password": "incorrect password",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, getErr)
}
