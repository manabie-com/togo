package sqllite

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

func TestRetrieveUserTaskLimit(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `INSERT INTO users (id, password, max_todo) VALUES('testUser', 'testPassword', 5);`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	eUserTaskLimit := 5

	// Call function to be tested
	aUserTaskLimit, err := dbHandler.RetrieveUserTaskLimit(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eUserTaskLimit)
	}

	if aUserTaskLimit != eUserTaskLimit {
		t.Log("Actual value:", aUserTaskLimit)
		t.Fatal("Expected value:", eUserTaskLimit)
	}
}

func TestRetrieveUserTaskLimitNoValue(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	eUserTaskLimit := 0

	// Call function to be tested
	aUserTaskLimit, err := dbHandler.RetrieveUserTaskLimit(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eUserTaskLimit)
	}

	if aUserTaskLimit != eUserTaskLimit {
		t.Log("Actual value:", aUserTaskLimit)
		t.Fatal("Expected value:", eUserTaskLimit)
	}
}

func TestRetrieveTasks(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `
		INSERT INTO tasks (id, content, user_id, created_date) VALUES(
			'25363527-604b-4cf8-b9ab-6e11336643dc',
			'first content',
			'testUser',
			'2021-10-15'
		);
		INSERT INTO tasks (id, content, user_id, created_date) VALUES(
			'81fb5603-e3de-4d80-b0b7-d863bb9779db',
			'second content',
			'testUser',
			'2021-10-15'
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testCreatedDate := "2021-10-15"

	var eTaskList []*storages.Task
	eTaskList = append(
		eTaskList,
		&storages.Task{
			ID:          "25363527-604b-4cf8-b9ab-6e11336643dc",
			Content:     "first content",
			UserID:      "testUser",
			CreatedDate: "2021-10-15",
		},
	)
	eTaskList = append(
		eTaskList,
		&storages.Task{
			ID:          "81fb5603-e3de-4d80-b0b7-d863bb9779db",
			Content:     "second content",
			UserID:      "testUser",
			CreatedDate: "2021-10-15",
		},
	)

	// Call function to be tested
	aTaskList, err := dbHandler.RetrieveTasks(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testCreatedDate,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Log("Expected value:")
		for _, task := range eTaskList {
			t.Logf("%+v", task)
		}
		t.Fatalf("Actual value is different from expected value.")
	}

	if !reflect.DeepEqual(aTaskList, eTaskList) {
		t.Log("Actual value:")
		for _, task := range aTaskList {
			t.Logf("%+v", task)
		}
		t.Log("Expected value:")
		for _, task := range eTaskList {
			t.Logf("%+v", task)
		}
		t.Fatalf("Actual value is different from expected value.")
	}
}

func TestRetrieveTasksNoValue(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testCreatedDate := "2021-10-15"

	var eTaskList []*storages.Task

	// Call function to be tested
	aTaskList, err := dbHandler.RetrieveTasks(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testCreatedDate,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Log("Expected value:")
		for _, task := range eTaskList {
			t.Logf("%+v", task)
		}
		t.Fatalf("Actual value is different from expected value.")
	}

	if !reflect.DeepEqual(aTaskList, eTaskList) {
		t.Log("Actual value:")
		for _, task := range aTaskList {
			t.Logf("%+v", task)
		}
		t.Log("Expected value:")
		for _, task := range eTaskList {
			t.Logf("%+v", task)
		}
		t.Fatalf("Actual value is different from expected value.")
	}
}

func TestCountTasks(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `
		INSERT INTO tasks (id, content, user_id, created_date) VALUES(
			'25363527-604b-4cf8-b9ab-6e11336643dc',
			'first content',
			'testUser',
			'2021-10-15'
		);
		INSERT INTO tasks (id, content, user_id, created_date) VALUES(
			'81fb5603-e3de-4d80-b0b7-d863bb9779db',
			'second content',
			'testUser',
			'2021-10-15'
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testCreatedDate := "2021-10-15"
	eTaskCount := 2

	// Call function to be tested
	aTaskCount, err := dbHandler.CountTasks(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testCreatedDate,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eTaskCount)
	}

	if aTaskCount != eTaskCount {
		t.Log("Actual value:", aTaskCount)
		t.Fatal("Expected value:", eTaskCount)
	}
}

func TestCountTasksNoValue(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testCreatedDate := "2021-10-15"
	eTaskCount := 0

	// Call function to be tested
	aTaskCount, err := dbHandler.CountTasks(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testCreatedDate,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eTaskCount)
	}

	if aTaskCount != eTaskCount {
		t.Log("Actual value:", aTaskCount)
		t.Fatal("Expected value:", eTaskCount)
	}
}

func TestAddTask(t *testing.T) {
	// Create in-memory users DB for mock values
	usersDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer usersDB.Close()
	usersDBHandler := NewLiteDB(usersDB)
	usersDBContext := context.Background()

	// Setup Mock DB
	usersDBStmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	usersDBHandler.DB.ExecContext(usersDBContext, usersDBStmt)
	usersDBStmt = `INSERT INTO users (id, password, max_todo) VALUES('testUser', 'testPassword', 5);`
	usersDBHandler.DB.ExecContext(usersDBContext, usersDBStmt)

	// Create in-memory tasks DB for mock values
	tasksDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer tasksDB.Close()
	tasksDBHandler := NewLiteDB(tasksDB)
	tasksDBContext := context.Background()

	// Setup Mock DB
	tasksDBStmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	tasksDBHandler.DB.ExecContext(tasksDBContext, tasksDBStmt)

	// Initialize other test values
	testTask := &storages.Task{
		ID:          "25363527-604b-4cf8-b9ab-6e11336643dc",
		Content:     "testContent",
		UserID:      "testUser",
		CreatedDate: "2021-10-15",
	}

	// Call function to be tested
	aReturn := tasksDBHandler.AddTask(
		tasksDBContext,
		testTask,
	)

	// Compare actual and expected values
	if aReturn != nil {
		t.Log("Actual result: ERROR")
		t.Fatal("Expected result: NO ERROR")
	}
}

func TestAddTaskPKViolation(t *testing.T) {
	// Create in-memory tasks DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `
		INSERT INTO tasks (id, content, user_id, created_date)
		VALUES(
			'25363527-604b-4cf8-b9ab-6e11336643dc',
			'testContent',
			'testUser',
			'2021-10-15'
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testTask := &storages.Task{
		ID:          "25363527-604b-4cf8-b9ab-6e11336643dc",
		Content:     "testContent",
		UserID:      "testUser",
		CreatedDate: "2021-10-15",
	}

	// Call function to be tested
	aReturn := dbHandler.AddTask(
		context,
		testTask,
	)

	// Compare actual and expected values
	if aReturn == nil {
		t.Log("Actual result: NO ERROR")
		t.Fatal("Expected result: ERROR")
	}
}

func TestAddTaskFKViolation(t *testing.T) {
	// Create in-memory users DB for mock values
	usersDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer usersDB.Close()
	usersDBHandler := NewLiteDB(usersDB)
	usersDBContext := context.Background()

	// Setup Mock DB
	usersDBStmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	usersDBHandler.DB.ExecContext(usersDBContext, usersDBStmt)
	usersDBStmt = `INSERT INTO users (id, password, max_todo) VALUES('testUser', 'testPassword', 5);`
	usersDBHandler.DB.ExecContext(usersDBContext, usersDBStmt)

	// Create in-memory tasks DB for mock values
	tasksDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer tasksDB.Close()
	tasksDBHandler := NewLiteDB(tasksDB)
	tasksDBContext := context.Background()

	// Setup Mock DB
	tasksDBStmt := `
		CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	tasksDBHandler.DB.ExecContext(tasksDBContext, tasksDBStmt)

	// Initialize other test values
	testTask := &storages.Task{
		ID:          "25363527-604b-4cf8-b9ab-6e11336643dc",
		Content:     "testContent",
		UserID:      "invalidTestUser",
		CreatedDate: "2021-10-15",
	}

	// Call function to be tested
	aReturn := tasksDBHandler.AddTask(
		tasksDBContext,
		testTask,
	)

	// Compare actual and expected values
	if aReturn == nil {
		t.Log("Actual result: NO ERROR")
		t.Fatal("Expected result: ERROR")
	}
}

func TestValidateUser(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `INSERT INTO users (id, password, max_todo) VALUES('testUser', 'testPassword', 5);`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testUserPassword := "testPassword"
	eValidation := true

	// Call function to be tested
	aValidation := dbHandler.ValidateUser(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testUserPassword,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eValidation)
	}

	if aValidation != eValidation {
		t.Log("Actual value:", aValidation)
		t.Fatal("Expected value:", eValidation)
	}
}

func TestValidateUserValueMismatch(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)
	stmt = `INSERT INTO users (id, password, max_todo) VALUES('testUser', 'testPassword', 5);`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testUserPassword := "wrongTestPassword"
	eValidation := false

	// Call function to be tested
	aValidation := dbHandler.ValidateUser(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testUserPassword,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eValidation)
	}

	if aValidation != eValidation {
		t.Log("Actual value:", aValidation)
		t.Fatal("Expected value:", eValidation)
	}
}

func TestValidateUserNoValue(t *testing.T) {
	// Create in-memory DB for mock values
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error opening DB", err)
	}
	defer db.Close()
	dbHandler := NewLiteDB(db)
	context := context.Background()

	// Setup Mock DB
	stmt := `
		CREATE TABLE users (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		);
	`
	dbHandler.DB.ExecContext(context, stmt)

	// Initialize other test values
	testUserID := "testUser"
	testUserPassword := "testPassword"
	eValidation := false

	// Call function to be tested
	aValidation := dbHandler.ValidateUser(
		context,
		sql.NullString{
			String: testUserID,
			Valid:  true,
		},
		sql.NullString{
			String: testUserPassword,
			Valid:  true,
		},
	)

	// Compare actual and expected values
	if err != nil {
		t.Log("Actual value:", err)
		t.Fatal("Expected value:", eValidation)
	}

	if aValidation != eValidation {
		t.Log("Actual value:", aValidation)
		t.Fatal("Expected value:", eValidation)
	}
}
