package model

import (
	"context"
	"log"
	"togo/app"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var Db = new(pg.DB)

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Task)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ConnectPGDatabase() {
	opt, err := pg.ParseURL(app.POSTGRESQL_URL)
	if err != nil {
		log.Fatal(err)
	}
	Db = pg.Connect(opt)
	err = Db.Ping(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = createSchema(Db)
	if err != nil {
		log.Fatal(err)
	}
}

func Initialize() {
	ConnectPGDatabase()
}

func ClosePGConnection() {
	Db.Close()
}
