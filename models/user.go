package models

import (
	"database/sql"
)

type User struct {
	ID            uint32 `json:"id" validate:"omitempty"`
	Email         string `json:"email" validate:"required,email,min=10,max=30"`
	Name          string `json:"name" validate:"required,min=5,max=20"`
	Password      string `json:"password" validate:"required,min=6,max=20"`
	IsPayment     bool   `json:"isPayment" validate:"omitempty"`
	IsActive      bool   `json:"isActive"`
	LimitDayTasks uint   `json:"limitDayTasks" validate:"omitempty"`
}

func (u *User) InsertUser(db *sql.DB) error {
	err := db.QueryRow(`INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id, name, email`, u.Name, u.Email, u.Password).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserById(db *sql.DB) error {
	err := db.QueryRow(`SELECT id, name, email, password, is_payment, is_active, limit_day_tasks FROM users WHERE id = $1`, u.ID).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsPayment, &u.IsActive, &u.LimitDayTasks)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByEmail(db *sql.DB) error {
	err := db.QueryRow(`SELECT id, name, email, password, is_payment, is_active, limit_day_tasks FROM users WHERE email = $1`, u.Email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsPayment, &u.IsActive, &u.LimitDayTasks)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ActiveUser(db *sql.DB) error {
	_, err := db.Exec(`UPDATE users SET is_active = $1 WHERE id = $2`, true, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec(`UPDATE users SET name = $1, email = $2 WHERE id = $3`, u.Name, u.Email, u.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser disable user active status => set is_active to false
func (u *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec(`UPDATE users SET is_active = $1 WHERE id = $2`, u.IsActive, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpgradePremium(db *sql.DB) error {
	err := db.QueryRow(`UPDATE users SET is_payment = $1, limit_day_tasks = $2 WHERE id = $3 RETURNING name, email`, u.IsPayment, u.LimitDayTasks, u.ID).Scan(&u.Name, &u.Email)
	if err != nil {
		return err
	}
	return nil
}
