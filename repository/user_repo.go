package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/qgdomingo/todo-app/model"
)

// Struct where the database connection memory address is stored
type UserRepository struct {
	DBPoolConn *pgxpool.Pool
}

// Function that takes the UserLogin model that should contain username and password 
// 	  then fetches the same username and password from the databse then compares the two.
// The compare can happen on the select query but I chose to separate it for a "future feature" 
// 	   where password encryption/decryption is implemented
func (u *UserRepository) LoginUserDB (user *model.UserLogin) (bool, map[string]string) {
	var userResult model.UserLogin
	message := make(map[string]string)
	sql := "SELECT username, password FROM users WHERE username = $1"

	rows, err := u.DBPoolConn.Query(context.Background(), sql, user.Username)

	if err != nil {
		message["message"] = "Failed to fetch user information from the users table"
		message["error"] = err.Error()
		return false, message
	}

	defer rows.Close()

	if rows.Next() {
		errRows := rows.Scan(&userResult.Username, &userResult.Password)

		if errRows != nil {
			message["message"] = "Error encountered when row data is being fetched"
			message["error"] = err.Error()
			return false, message
		}

		if userResult.Username == user.Username && userResult.Password == user.Password {
			return true, nil
		}
		
	}

	return false, nil
}

// Function that creates a new user and if created succesfully, will create a task limit for the user immediately
func (u *UserRepository) RegisterUserDB (user *model.NewUser) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "INSERT INTO users (username, name, email, password) VALUES ($1, $2, $3, $4) RETURNING username"

	rows, err := u.DBPoolConn.Query(context.Background(), sql, user.Username, user.Name, user.Email, user.Password)

	if err != nil {
		message["message"] = "Failed to insert new user to the users table"
		message["error"] = err.Error()
		return false, message
	}

	defer rows.Close()

	if rows.Next() {
		sql = "INSERT INTO task_limit_config (username, task_limit) VALUES ($1, $2) RETURNING username"

		rowsConfig, errConfig := u.DBPoolConn.Query(context.Background(), sql, user.Username, user.TaskLimit)

		if errConfig != nil {
			message["message"] = "Failed to insert new user to the task limit config table"
			message["error"] = err.Error()
			return false, message
		}

		defer rowsConfig.Close()

		if rowsConfig.Next() {
			return true, nil
		}
	} 

	return false, nil
}

// Function that fetches user details and as well as the configured task limit for the same user
func (u *UserRepository) FetchUserDetailsDB (username string) ([]model.UserDetails, map[string]string) {
	var userList []model.UserDetails
	message := make(map[string]string)
	sql := "SELECT u.username, u.name, u.email, tlc.task_limit FROM users u, task_limit_config tlc WHERE tlc.username = u.username AND u.username = $1"

	rows, err := u.DBPoolConn.Query(context.Background(), sql, username)

	if err != nil {
		message["message"] = "Failed to fetch user information from the users table"
		message["error"] = err.Error()
		return nil, message
	}

	defer rows.Close()

	if rows.Next() {
		var user model.UserDetails
		err = rows.Scan(&user.Username, &user.Name, &user.Email, &user.TaskLimit)

		if err != nil {
			message["message"] = "Error encountered when row data is being fetched"
			message["error"] = err.Error()
			return nil, message
		}

		userList = append(userList, user)
	}

	return userList, nil
}

// Function that updates the user details and if successful, will update the configured task limit for the same user
func (u *UserRepository) UpdateUserDetailsDB (user *model.UserDetails, username string) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "UPDATE users SET username = $1, name = $2, email = $3 WHERE username = $4 RETURNING username"

	rows, err := u.DBPoolConn.Query(context.Background(), sql, user.Username, user.Name, user.Email, username)

	if err != nil {
		message["message"] = "Unable to update data into the users table"
		message["error"] = err.Error()
		return false, message
	}

	defer rows.Close()

	if rows.Next() {
		sql = "UPDATE task_limit_config SET username = $1, task_limit = $2 WHERE username = $3 RETURNING username"

		rowsConfig, errConfig := u.DBPoolConn.Query(context.Background(), sql, user.Username, user.TaskLimit, username)

		if errConfig != nil {
			message["message"] = "Failed to update user to the task limit config table"
			message["error"] = err.Error()
			return false, message
		}

		defer rowsConfig.Close()

		if rowsConfig.Next() {
			return true, nil
		}
	} 

	return false, nil
}

// Function that changes the user's password, this takes the current password of the user and compares it in the database
// 		This is the opposite of the implementation in LoginUserDB function
func (u *UserRepository) UserPasswordChangeDB (user *model.UserNewPassword, username string) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "UPDATE users SET password = $1 WHERE password = $2 AND username = $3 RETURNING username"

	row, err := u.DBPoolConn.Query(context.Background(), sql, user.NewPassword, user.CurrentPassword, username)

	if err != nil {
		message["message"] = "Unable to update data into the users table"
		message["error"] = err.Error()
		return false, message
	}

	defer row.Close()

	return row.Next(), nil
}