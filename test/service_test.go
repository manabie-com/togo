package test

import (
	"database/sql"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
	"log"
	"testing"
)

func TestGetNonExistentProduct(t *testing.T) {
	//db, dB := SetupTests1()
	////var name string
	//user := model.Users{ID: "firstUser", Password: "example", MaxTodo: "5"}
	//commonReply := []map[string]interface{}{{"id": "firstUser", "password": "example", "max_todo": "30"}}
	////mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO users").WithArgs(user)
	//mocket.Catcher.Reset().NewMock().WithQuery("Select * from users").WithReply(commonReply)
	////dB.Exec("")
	//var userTest = model.Users{}
	//err := dB.QueryRow("Select * from users").Scan(&userTest.ID, &userTest.Password, &userTest.MaxTodo)
	//
	//
	////rows, err := db.Exec("Select * from users").Rows()
	//if err != nil {}
	////user := model.Users{ID: "firstUser", Password: "example", MaxTodo: "5"}
	////var mockedId int64 = 64
	//
	//
	////var userTest model.Users
	////db.AutoMigrate(&model.Users{})
	//result := db.Create(user).Commit()
	//
	//print(result.Error)
	//print(result.RowsAffected)
	//db.Table("users").Row()
	//print(userTest.ID)

	//if assert.NotNil(t, resp) {
	//	assert.Equal(t, uint32(mockedId), resp.Group.Id)
	//	assert.Equal(t, name, resp.Group.Name)
	//	assert.Equal(t, description, resp.Group.Description)
	//}

}

//func provideMock() (sqlmock.Sqlmock, *gorm.DB,) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		print("asdasd")
//		//t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectBegin()
//	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//
//	// now we execute our method
//	if err = recordStats(db, 2, 3); err != nil {
//		print("asdasd")
//		//t.Errorf("error was not expected while updating stats: %s", err)
//	}
//
//	// we make sure that all expectations were met
//	if err := mock.ExpectationsWereMet(); err != nil {
//
//		//t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//	// open the database for gorm (*gorm.DB)
//	user := model.Users{ID: "firstUser", Password: "example", MaxTodo: "5"}
//
//	mockedGorm, _ := gorm.Open("sqlmock", "sqlmock_db_0")
//	mockedGorm.AutoMigrate(&model.Users{})
//	mockedGorm.Create(&user)
//	result := mockedGorm.Create(&user)
//	print(result.Error)
//	print(result.RowsAffected)
//	var userTest model.Users
//	mockedGorm.Where("id = ? AND password = ?", "firstUser", "example").Find(&userTest)
//
//	return mock, mockedGorm
//}


func SetupTests() *gorm.DB {
	user := model.Users{ID: "firstUser", Password: "example", MaxTodo: "5"}
	commonReply := []map[string]interface{}{{"id": "firstUser", "password": "example", "max_todo": "30"}}

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO users").WithArgs(user)
	mocket.Catcher.Reset().NewMock().WithQuery("SELECT * FROM users").WithArgs(user).WithReply(commonReply)

	db, err := gorm.Open(mocket.DriverName, "connection_string")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}
	//db, err := sql.Open(mocket.DriverName, "connection_string")

	// Log mode shows the query gorm uses, so we can replicate and mock it
	db.LogMode(true)

	return db
}

func SetupTests1() (*gorm.DB, *sql.DB) { // or *gorm.DB
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	// GORM
	db, err := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string

	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}

	db1, err1 := sql.Open(mocket.DriverName, "connection_string") // Can be any connection string

	if err1 != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}

	// OR
	// Regular sql package usage

	return db, db1
}