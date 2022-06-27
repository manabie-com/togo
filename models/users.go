package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
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

func (r *Repository)GetAllUser() ([]User, error) { // Get all user from the database
	rows, err := r.DB.Query("SELECT * FROM users;")
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

func (r *Repository)InsertUser(user NewUser) error { // Insert one user to the database
	if !CheckUserInput(user) {
		return errors.New("decode failed")
	}
	_, err := r.DB.Exec("INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3);", user.Username, user.Password, user.LimitTask)
	return err
}

func (r *Repository)DeleteUser(id int) error { // delete 1 user
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *Repository)UpdateUser(newUser NewUser, id int) error { // Update one user already exist in database
	if !CheckUserInput(newUser) {
		return errors.New("user input invalid")
	}
	_, err := r.DB.Exec("UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $4", newUser.Username, newUser.Password, newUser.LimitTask, id)
	return err
}

func (r *Repository)FindUserByID(id int) (User, bool) { // Check ID is valid or not
	user := User{}
	row := r.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return user, false
	}
	return user, true
}

func CheckUserInput(user NewUser) bool { // Check user input is valid or not
	password := strings.TrimSpace(user.Password)
	username := strings.TrimSpace(user.Username)
	if password == "" || username == ""{
		return false
	}
	return true
}

func (r *Repository)CheckUserNameExist(username string) (User, bool) { // Check if username already exist or not
	user := User{}
	row := r.DB.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.LimitTask)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Error checking if row exist")
		}
		return user, false
	}
	return user, true
}
