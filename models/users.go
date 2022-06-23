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
}
type NewUser struct {
	Username string
	Password string
}

func GetAllUser() ([]User, error) { // Get all user from the database
	rows, err := DB.Query("SELECT * FROM users;")
	var users []User
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	return users, nil
}


func InsertUser(user NewUser) error { // Insert one user to the database
	if !CheckUser(user) {
		return errors.New("decode failed")
	}
	_, err := DB.Exec("INSERT INTO users(username, password) VALUES ($1, $2);", user.Username, user.Password)
	if err != nil {
		return errors.New("insert database failed")
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func UpdateUser(newUser NewUser, id int) error { // Update one user already exist in database
	if !CheckUser(newUser) {
		return errors.New("user invalid")
	}
	_, err := DB.Exec("UPDATE users SET username = $1, password = $2 WHERE id = $3", newUser.Username, newUser.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func CheckIDAndReturn(id int) (User, bool) { // Check ID is valid or not
	user := User{}

	row := DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
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
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist")
		}
		return user, false
	}

	return user, true
}