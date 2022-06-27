package repository

import (
	e "lntvan166/togo/internal/entities"
)

// CREATE

func (r *repository) AddUser(u *e.User) error {
	const query = `INSERT INTO users (username, password, plan, max_todo) VALUES ($1, $2, $3, $4)`
	tx, err := r.DB.Begin()
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

// READ

func (r *repository) GetAllUsers() ([]*e.User, error) {
	const query = `SELECT * FROM users`
	rows, err := r.DB.Query(query)
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

func (r *repository) GetUserByName(username string) (*e.User, error) {
	const query = `SELECT * FROM users WHERE username = $1`
	u := &e.User{}
	err := r.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repository) GetUserByID(id int) (*e.User, error) {
	const query = `SELECT * FROM users WHERE id = $1`
	u := &e.User{}
	err := r.DB.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repository) GetNumberOfTaskTodayByUserID(id int) (int, error) {
	const query = `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND DATE(created_at) = CURRENT_DATE`
	var count int
	err := r.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) GetPlanByID(id int) (string, error) {
	const query = `SELECT plan FROM users WHERE id = $1`
	var plan string
	err := r.DB.QueryRow(query, id).Scan(&plan)
	if err != nil {
		return "", err
	}
	return plan, nil
}

func (r *repository) GetPlanByUsername(username string) (string, error) {
	const query = `SELECT plan FROM users WHERE username = $1`
	var plan string
	err := r.DB.QueryRow(query, username).Scan(&plan)
	if err != nil {
		return "", err
	}
	return plan, nil
}

// func (r *repository) CheckUserExist(username string) (bool, error) {
// 	const query = `SELECT * FROM users WHERE username = $1`
// 	u := &e.User{}
// 	err := r.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Plan, &u.MaxTodo)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// UPDATE

func (r *repository) UpdateUser(u *e.User) error {
	const query = `UPDATE users SET username = $1, password = $2, plan = $3, max_todo = $4 WHERE id = $5`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, u.Username, u.Password, u.Plan, u.MaxTodo, u.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repository) UpgradePlan(id int, plan string, limit int) error {
	const query = `UPDATE users SET plan = $1, max_todo = $2 WHERE id = $3`
	tx, err := r.DB.Begin()
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

// DELETE

func (r *repository) DeleteUserByID(id int) error {
	const query = `DELETE FROM users WHERE id = $1`
	tx, err := r.DB.Begin()
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
