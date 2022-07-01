package models

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	LimitTask int
}
type NewUser struct {
	Username  string
	Password  string
	LimitTask int
}

// Get all user from the database
func (Conn *DbConn) GetAllUser() ([]User, error) {
	rows, err := Conn.DB.Query(QueryAllUserText)
	var users []User
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Insert one user to the database
func (Conn *DbConn) InsertUser(user NewUser) error {
	if !CheckUserInput(user) {
		return errors.New("decode failed")
	}
	_, err := Conn.DB.Exec(InsertUserText, user.Username, user.Password, user.LimitTask)
	return err
}

// delete 1 user
func (Conn *DbConn) DeleteUser(id int) error {
	if id == 1 {
		return errors.New("can't delete admin")
	}
	_, err := Conn.DB.Exec(DeleteUserText, id)
	return err
}

// Update one user already exist in database
func (Conn *DbConn) UpdateUser(newUser NewUser, id int) error {
	if id ==1 {
		return errors.New("can't update admin account")
	}
	if !CheckUserInput(newUser) {
		return errors.New("user input invalid")
	}
	_, err := Conn.DB.Exec(UpdateUserText, newUser.Username, newUser.Password, newUser.LimitTask, id)
	return err
}

// Check ID is valid or not
func (Conn *DbConn) FindUserByID(id int) (User, bool) {
	user := User{}
	row := Conn.DB.QueryRow(FindUserByIDText, id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return user, false
	}
	return user, true
}

// Check if username already exist or not
func (Conn *DbConn) CheckUserNameExist(username string) (User, bool) {
	user := User{}
	row := Conn.DB.QueryRow(QueryAllUsernameText, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist, err: " + err.Error())
		}
		return user, false
	}
	return user, true
}
