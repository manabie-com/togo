package model

import (
	"lntvan166/togo/db"
	e "lntvan166/togo/entities"
)

func AddUser(u *e.User) error {
	const query = `INSERT INTO users (username, password, plan, max_todo) VALUES ($1, $2, $3, $4)`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, u.Username, u.Password, u.Plan, u.MaxTodo)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UpdateUser(u *e.User) error {
	const query = `UPDATE users SET username = $1, password = $2, plan = $3, max_todo = $4 WHERE id = $5`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, u.Password, u.Plan, u.MaxTodo, u.Username)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteUser(id int) error {
	const query = `DELETE FROM users WHERE id = $1`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetUserByName(username string) (*e.User, error) {
	const query = `SELECT * FROM users WHERE username = $1`
	u := &e.User{}
	err := db.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetAllUsers() ([]*e.User, error) {
	const query = `SELECT * FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*e.User{}
	for rows.Next() {
		u := &e.User{}
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func CheckUserExist(username string) bool {
	const query = `SELECT * FROM users WHERE username = $1`
	u := &e.User{}
	err := db.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
	return err == nil
}

func GetUserIDByUsername(username string) (int, error) {
	const query = `SELECT id FROM users WHERE username = $1`
	var id int
	err := db.DB.QueryRow(query, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetMaxTaskByUserID(id int) (int, error) {
	const query = `SELECT max_todo FROM users WHERE id = $1`
	var max int
	err := db.DB.QueryRow(query, id).Scan(&max)
	if err != nil {
		return 0, err
	}
	return max, nil
}

func GetPlanByID(id int) (string, error) {
	const query = `SELECT plan FROM users WHERE id = $1`
	var plan string
	err := db.DB.QueryRow(query, id).Scan(&plan)
	if err != nil {
		return "", err
	}
	return plan, nil
}

func GetPlanByUsername(username string) (string, error) {
	const query = `SELECT plan FROM users WHERE username = $1`
	var plan string
	err := db.DB.QueryRow(query, username).Scan(&plan)
	if err != nil {
		return "", err
	}
	return plan, nil
}

func UpgradePlan(id int, plan string, limit int) error {
	const query = `UPDATE users SET plan = $1, max_todo = $2 WHERE id = $3`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, plan, limit, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetNumberOfTaskToday(id int) (int, error) {
	const query = `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND DATE(created_at) = CURRENT_DATE`
	var count int
	err := db.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
