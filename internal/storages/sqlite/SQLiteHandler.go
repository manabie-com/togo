package sqllite

import (
	// "context"
	"fmt"
	"database/sql"
	"log"
	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type SQLiteHandler struct {
	DB *sql.DB
}

func (handler *SQLiteHandler) Execute(statement string) error {
	_, err := handler.DB.Exec(statement)
	return err
}

func (handler *SQLiteHandler) Query(statement string) (storages.IRow, error) {
	rows, err := handler.DB.Query(statement)
	if err != nil {
		log.Println("query failed")
		fmt.Println(err)
		return new(LiteRow), err
	}

	row := new(LiteRow)
	row.Rows = rows	
	
	return row, nil
}

type LiteRow struct {
	Rows *sql.Rows
}

func (r LiteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		log.Println("Scan error!")
		return err
	}
	return nil
}

func (r LiteRow) Next() bool {
	return r.Rows.Next()
}