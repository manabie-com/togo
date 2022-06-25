package models

import (
	"database/sql"
)

type User struct {
	ID            uint32 `json:"id" validate:"omitempty"`
	Email         string `json:"email" validate:"required,email"`
	Name          string `json:"name" validate:"required,min=5,max=20"`
	Password      string `json:"password" validate:"required,min=6,max=20"`
	IsPayment     bool   `json:"isPayment" validate:"omitempty"`
	IsActive      bool   `json:"isActive"`
	LimitDayTasks uint   `json:"limitDayTasks" validate:"omitempty"`
}

func (u *User) InsertOne(db *sql.DB) error {
	err := db.QueryRow(`INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id, name, email`, u.Name, u.Email, u.Password).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetOneById(db *sql.DB) error {
	err := db.QueryRow(`SELECT id, name, email, password, is_payment, is_active, limit_day_tasks FROM users WHERE id = $1`, u.ID).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsPayment, &u.IsActive, &u.LimitDayTasks)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetOneByEmail(db *sql.DB) error {
	err := db.QueryRow(`SELECT id, name, email, password, is_payment, is_active, limit_day_tasks FROM users WHERE email = $1`, u.Email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsPayment, &u.IsActive, &u.LimitDayTasks)
	if err != nil {
		return err
	}
	return nil
}
