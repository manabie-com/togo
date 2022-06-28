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
func (Conn *DbConn)GetAllUser() ([]User, error) {
	rows, err := Conn.DB.Query("SELECT * FROM users")
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
func (Conn *DbConn)InsertUser(user NewUser) error {
	if !CheckUserInput(user) {
		return errors.New("decode failed")
	}
	_, err := Conn.DB.Exec("INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3)", user.Username, user.Password, user.LimitTask)
	return err
}

// delete 1 user
func (Conn *DbConn)DeleteUser(id int) error {
	_, err := Conn.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

// Update one user already exist in database
func (Conn *DbConn)UpdateUser(newUser NewUser, id int) error {
	if !CheckUserInput(newUser) {
		return errors.New("user input invalid")
	}
	_, err := Conn.DB.Exec("UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $4", newUser.Username, newUser.Password, newUser.LimitTask, id)
	return err
}

// Check ID is valid or not
func (Conn *DbConn)FindUserByID(id int) (User, bool) {
	user := User{}
	row := Conn.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
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
func (Conn *DbConn)CheckUserNameExist(username string) (User, bool) {
	user := User{}
	row := Conn.DB.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist, err: "+ err.Error())
		}
		return user, false
	}
	return user, true
}
