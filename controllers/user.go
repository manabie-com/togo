package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var SignUp = func(db *sql.DB, w http.ResponseWriter, r *http.Request) interface{} {
	defer db.Close()
	user := &models.User{
		IsPayment:     false,
		LimitDayTasks: 10,
	}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid request, Please check your input fields!", nil)
		return err
	}

	results, err := db.Exec(`INSERT INTO users(name, email, password) VALUES($1,$2,$3) RETURNING id`, user.Name, user.Email, user.Password)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "OOPS, something were wrong, please try again", nil)
		return err
	}
	// send token jwt here
	// ...
	u.Respond(w, http.StatusCreated, "Success", "Created Account", results)
	return nil
}

// var Login = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {

// 	// get email password
// 	// if email and password not exist
// 	u.Respond(w, http.StatusBadRequest, map[string]interface{}{})
// 	// if email exist and password incorrect
// 	u.Respond(w, http.StatusUnauthorized, map[string]interface{}{})
// 	// if email and password OK
// 	// send token to client
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }

var GetMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode jwt => userId

	// data := db.QueryRow("SELECT name, email, is_payment, limit_day_tasks FROM users WHERE id = $1", 1)
	var id int = 1
	rows, err := db.Query(`SELECT name, email, is_payment FROM users WHERE id = $1`, id)
	if err != nil {
		fmt.Println(err)
	}

	var user models.User

	for rows.Next() {
		var (
			id    int = 1
			name  string
			email string
		)
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
		// check errors
		user = models.User{Name: name, Email: email}
		fmt.Println(user)
	}

	// var user User

	u.Respond(w, http.StatusOK, "Success", "Success", user)
	// if jwt  invalid
	// u.Respond(w, http.StatusBadRequest, , )
	// decode jwt -> get userID -> get user
	// u.Respond(w, http.StatusOK, map[string]interface{}{})
}

// var UpdateMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// success
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }
