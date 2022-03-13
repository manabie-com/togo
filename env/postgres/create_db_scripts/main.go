package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/HoangMV/togo/lib/pgsql"
	"github.com/spf13/viper"
)

// ================= MAIN =====================
func init() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	pgsql.Install()

	for {
		if err := Run(); err != nil {
			fmt.Printf("can not to apply migrations, err: %+v\n", err)
			continue
		}

		break
	}
}

// ================= RUN =====================

// Run: create tables or updates
//
// @return: error
func Run() error {
	if ok, err := tableExist(vers); err != nil {
		return err
	} else if !ok {
		ctx, cancel := pgsql.GetDefaultContext()
		defer cancel()

		if _, err := pgsql.Get().ExecContext(ctx, createTables[vers]); err != nil {
			return err
		}
	}

	mapRunErrors := make(map[string]string)
	mapTmp := createTables
	for id, sqlStr := range createTables {
		if id == vers {
			continue
		}

		if ok, err := migrationExist(id); err != nil {
			return fmt.Errorf("can not check migration exist, err: %+v", err)
		} else if ok {
			delete(mapTmp, id)
			continue
		}

		ctx, cancel := pgsql.GetDefaultContext()
		defer cancel()

		fmt.Println("Id: ", id, " SqlString: ", sqlStr)
		tx, err := pgsql.Get().BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("can not start transaction for migration: %s, err: %+v", id, err)
		}
		defer tx.Rollback()

		if _, err := tx.ExecContext(ctx, sqlStr); err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("can not execute migration: %s, err: %+v", id, err)
				if err := applyVersionWitouttransaction(id); err != nil {
					log.Printf("applyVersionWitouttransaction: %s, err: %+v", id, err)
				}
			} else {
				errTr := tx.Rollback()
				mapRunErrors[id] = fmt.Sprintf("can not execute migration: %s, err: %+v, transaction: %+v", id, err, errTr)
				continue
			}
		} else if err := applyVersion(tx, id); err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("can not execute migration: %s, err: %+v", id, err)
			} else {
				errTr := tx.Rollback()
				mapRunErrors[id] = fmt.Sprintf("can not apply migration: %s, err: %+v, transaction: %+v", id, err, errTr)
				continue
			}
		} else if err := tx.Commit(); err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("can not execute migration: %s, err: %+v", id, err)
			} else {
				errTr := tx.Rollback()
				mapRunErrors[id] = fmt.Sprintf("can not comming transaction: %s, err: %+v, transaction: %+v", id, err, errTr)
				continue
			}
		}

		delete(mapTmp, id)
	}

	createTables = mapTmp

	if len(mapRunErrors) > 0 {
		log.Printf("Migration runner errors: %+v", mapRunErrors)
		return fmt.Errorf("migration runner errors")
	}

	return nil
}

func migrationExist(version string) (bool, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	var exist bool
	err := pgsql.Get().QueryRowContext(ctx, "select exists(select 1 from "+vers+" where id=$1)", version).Scan(&exist)

	return exist, err
}

func applyVersion(tx *sql.Tx, version string) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	_, err := tx.ExecContext(ctx, "INSERT INTO "+vers+" (id, created_at) VALUES ($1, $2)", version, time.Now())
	return err
}

func applyVersionWitouttransaction(version string) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	_, err := pgsql.Get().ExecContext(ctx, "INSERT INTO "+vers+" (id, created_at) VALUES ($1, $2)", version, time.Now())
	return err
}

func tableExist(name string) (bool, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	var exist bool
	err := pgsql.Get().QueryRowContext(ctx, `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1);`, name).Scan(&exist)

	return exist, err
}

// ================= SCRIPTs =====================

const vers = "db_version"

var createTables = map[string]string{
	vers: `
		CREATE TABLE ` + vers + ` (
			id VARCHAR(255) PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP     
		);
		`,

	"create_trigger_function": `
		CREATE OR REPLACE FUNCTION update_updated_at_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = now();
			RETURN NEW;
		END;
		$$ language 'plpgsql';
		`,

	"create_users_tables": `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY NOT NULL,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(128) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL  DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE UNIQUE INDEX users_idx_username
			ON users (username);
		
		CREATE TRIGGER tg_users_updated_at
    	BEFORE UPDATE
    	ON  users
    	FOR EACH ROW
    	EXECUTE PROCEDURE update_updated_at_column();
		
		`,
	"create_todo_tables": `
		CREATE TABLE todos (
			id SERIAL PRIMARY KEY NOT NULL,
			user_id INT NOT NULL,
			content TEXT NOT NULL,
			status INT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL  DEFAULT CURRENT_TIMESTAMP,
		
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		);

		CREATE TRIGGER tg_todos_updated_at
		BEFORE UPDATE
		ON  todos
		FOR EACH ROW
		EXECUTE PROCEDURE update_updated_at_column();


		CREATE TABLE user_todo_config (
			user_id INT PRIMARY KEY NOT NULL, 
			max_todo INT NOT NULL,

			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL  DEFAULT CURRENT_TIMESTAMP,
		
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		);

		CREATE TRIGGER tg_user_todo_config_updated_at
		BEFORE UPDATE
		ON  user_todo_config
		FOR EACH ROW
		EXECUTE PROCEDURE update_updated_at_column();
		`,
}
