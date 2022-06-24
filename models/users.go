package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type User struct {
	Id       int64
	Username string
	Password string
	LimitTask    int
}
type NewUser struct {
	Username string
	Password string
	LimitTask    int
}

func GetAllUser() ([]User, error) { // Get all user from the database
	rows, err := DB.Query("SELECT * FROM users;")
	var users []User
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask); err != nil{
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func InsertUser(user NewUser) error { // Insert one user to the database
	if !CheckUser(user) {
		return errors.New("decode failed")
	}
	_, err := DB.Exec("INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3);", user.Username, user.Password, user.LimitTask)
	return err
}

func DeleteUser(id int) error { // delete 1 user
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func UpdateUser(newUser NewUser, id int) error { // Update one user already exist in database
	if !CheckUser(newUser) {
		return errors.New("user invalid")
	}
	_, err := DB.Exec("UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $3", newUser.Username, newUser.Password, newUser.LimitTask,id)
	return err

}

func CheckIDUserAndReturn(id int) (User, bool) { // Check ID is valid or not
	user := User{}
	row := DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist")
		}
		return user, false
	}
	return user, true
}

func CheckUser(user NewUser) bool { // Check user input is valid or not
	password := strings.TrimSpace(user.Password)
	username := strings.TrimSpace(user.Username)
	if password == "" || username == "" {
		return false
	}
	return true
}

func CheckUserInput(username string) (User, bool) { // Check if username already exist or not
	user := User{}
	row := DB.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist")
		}
		return user, false
	}
	return user, true
}
