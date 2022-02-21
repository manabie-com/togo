package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"togo/domain"
	"togo/infrastructure/database/structs"
)

type DbGormStruct struct {
	*gorm.DB
}

var dbBuiltin *sql.DB
var dbGorm DbGormStruct

func Init(dbInfo string) *DbGormStruct {
	dbBuiltin, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic("failed to connect mysql database")
	}

	dbBuiltin.SetConnMaxLifetime(3 * time.Minute)
	dbBuiltin.SetMaxOpenConns(100)
	dbBuiltin.SetMaxIdleConns(100)

	dbGorm.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbBuiltin,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm database")
	}

	return &dbGorm
}

func (receiver DbGormStruct) Get(getStruct structs.GetStruct) structs.GetResultStruct {
	var getResultStruct structs.GetResultStruct
	model := initModelMapTable(getStruct.TableName)

	getResultStruct.TableName = getStruct.TableName
	getResultStruct.Error = nil
	getResultStruct.Status = "init"
	getResultStruct.Message = "init result"
	getResultStruct.Conditions = getStruct.Conditions
	getResultStruct.Data = types.Interface{}

	result := dbGorm.Table(getStruct.TableName).Where(getStruct.Conditions).Scan(&model)
	if result.Error != nil {
		getResultStruct.Error = result.Error
		getResultStruct.Status = "error"
		getResultStruct.Message = result.Error.Error()

		return getResultStruct
	}

	getResultStruct.Status = "success"
	getResultStruct.Message = "success result"
	getResultStruct.Data = model

	return getResultStruct
}

func initModelMapTable(tableName string) interface{} {
	switch tableName {
	case "users":
		return new(domain.User)
	case "todos":
		return new(domain.Todo)
	}

	return types.Interface{}
}
