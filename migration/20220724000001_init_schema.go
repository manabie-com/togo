package migration

import (
	"context"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/uptrace/bun"
)

type InitSchema struct {
	Migration
	Version string
}

func (m *InitSchema) Up() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {

		// create "users" table
		const users = `CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			is_active BOOLEAN DEFAULT false NOT NULL,
			daily_limit INT NOT NULL DEFAULT 3,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)`
		if _, err = tx.Exec(users); err != nil {
			return err
		}

		// create "tasks" table
		const tasks = `CREATE TABLE IF NOT EXISTS tasks(
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			published_date DATE NOT NULL,
			status SMALLINT NOT NULL DEFAULT 0,
			created_by INT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			CONSTRAINT fk_owner
      		FOREIGN KEY(created_by) 
	  			REFERENCES users(id)
	  				ON DELETE RESTRICT
		)`
		if _, err = tx.Exec(tasks); err != nil {
			return err
		}

		return nil
	})
}

func (m *InitSchema) Down() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS tasks, users`)
		return err
	})
}

func (m *InitSchema) GetVersion() string {
	return "20220724000001"
}
