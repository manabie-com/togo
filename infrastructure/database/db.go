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

func (receiver DbGormStruct) Create(createStruct structs.CreateStruct) structs.CreateResultStruct {
	var createResultStruct structs.CreateResultStruct

	createResultStruct.TableName = createStruct.TableName
	createResultStruct.Error = nil
	createResultStruct.Status = "init"
	createResultStruct.Message = "init result"
	createResultStruct.Data = types.Interface{}

	result := dbGorm.Table(createStruct.TableName).Create(createStruct.Data)
	if result.Error != nil {
		createResultStruct.Error = result.Error
		createResultStruct.Status = "error"
		createResultStruct.Message = result.Error.Error()

		return createResultStruct
	}

	createResultStruct.Status = "success"
	createResultStruct.Message = "success result"
	createResultStruct.Data = createStruct.Data

	return createResultStruct
}

func (receiver DbGormStruct) Get(getStruct structs.GetStruct) structs.GetResultStruct {
	var getResultStruct structs.GetResultStruct
	_, models := initModelMapTable(getStruct.TableName)

	getResultStruct.TableName = getStruct.TableName
	getResultStruct.Error = nil
	getResultStruct.Status = "init"
	getResultStruct.Message = "init result"
	getResultStruct.Conditions = getStruct.Conditions
	getResultStruct.Data.RowsAffected = 0
	getResultStruct.Data.Result = types.Interface{}

	result := dbGorm.Table(getStruct.TableName).Where(getStruct.Conditions).Find(models)

	if result.Error != nil {
		getResultStruct.Error = result.Error
		getResultStruct.Status = "error"
		getResultStruct.Message = result.Error.Error()

		return getResultStruct
	}

	getResultStruct.Status = "success"
	getResultStruct.Message = "success result"
	getResultStruct.Data.RowsAffected = result.RowsAffected
	getResultStruct.Data.Result = models

	return getResultStruct
}

func initModelMapTable(tableName string) (interface{}, interface{}) {
	switch tableName {
	case "users":
		return new(domain.User), new([]domain.User)
	case "todos":
		return new(domain.Todo), new([]domain.Todo)
	case "todo_limit":
		return new(domain.TodoLimit), new([]domain.TodoLimit)
	}
	return types.Interface{}, types.Interface{}
}
