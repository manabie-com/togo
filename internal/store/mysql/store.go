package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kozloz/togo"
)

const TimeFormat = "2006-01-02 15:04:05"

type Store struct {
	db *sql.DB
}

func NewStore(dbName string, host string, port string, user string, pass string) (*Store, error) {
	log.Println("Creating mysql store")
	login := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	db, err := sql.Open("mysql", login)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping db: %v", err)
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) CreateTask(userID int64, task string) (*togo.Task, error) {
	log.Println("Creating task in database")
	query := `
		INSERT INTO tasks (user_id, name) values(?,?)
	`
	res, err := s.db.Exec(query, userID, task)
	if err != nil {
		log.Printf("Failed to create task with error: '%v'.", err)
		return nil, err
	}
	id, _ := res.LastInsertId()

	return &togo.Task{
		ID:     id,
		UserID: userID,
		Name:   task,
	}, nil
}

func (s *Store) GetUserTasks(userID int64) ([]*togo.Task, error) {
	log.Printf("Getting tasks of user '%d'.", userID)
	query := `
		SELECT id, name from tasks where user_id = ?
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		log.Printf("Failed to query user tasks with error: '%v'", err)
		return nil, err
	}
	defer rows.Close()
	var tasks []*togo.Task
	for rows.Next() {
		task := &togo.Task{}
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			log.Printf("Failed to scan user task with error: '%v'", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Store) GetUser(userID int64) (*togo.User, error) {
	log.Printf("Getting user with id '%d'", userID)

	// Get user state
	user := &togo.User{}
	selectUserQuery := `
		SELECT id, daily_limit from users where id = ?
	`
	userRow := s.db.QueryRow(selectUserQuery, userID)
	err := userRow.Scan(&user.ID, &user.DailyLimit)
	if err == sql.ErrNoRows {
		log.Printf("No user found with id '%d'.", userID)
		return nil, nil
	}
	if err != nil {
		log.Printf("Failed to query users with error: '%v'.", err)
		return nil, err
	}

	log.Printf("Getting user counters")
	selectCounterQuery := `
		SELECT daily_count, last_updated from user_daily_counters where user_id = ?
	`

	// Get user tasks
	tasks, err := s.GetUserTasks(userID)
	if err != nil {
		return nil, err
	}
	user.Tasks = tasks

	// Get user daily counter
	counter := &togo.DailyCounter{}
	counterRow := s.db.QueryRow(selectCounterQuery, userID)
	lastUpdatedStr := ""
	err = counterRow.Scan(&counter.DailyCount, &lastUpdatedStr)
	if err == sql.ErrNoRows {
		log.Printf("No user counter found for user id '%d'.", userID)
		return user, nil
	}
	if err != nil {
		log.Printf("Failed to query user counters with error: '%v'.", err)
		return nil, err
	}
	counter.LastUpdated, _ = time.Parse(TimeFormat, lastUpdatedStr)
	user.DailyCounter = counter

	return user, nil
}
func (s *Store) CreateUser(userID int64) (*togo.User, error) {
	log.Println("Creating user in database")
	query := `
		INSERT INTO users (id, daily_limit) values(?,0)
	`
	res, err := s.db.Exec(query, userID)
	if err != nil {
		log.Printf("Failed to create user with error: '%v'.", err)
		return nil, err
	}
	id, _ := res.LastInsertId()

	return &togo.User{
		ID: id,
	}, nil
}

func (s *Store) UpdateUser(user *togo.User) (*togo.User, error) {
	log.Println("Updating user in database")

	// Update user object

	// Update counter only if user has one
	if user.DailyCounter == nil {
		return user, nil
	}

	log.Println("Updating user counter in database")
	query := `
		 INSERT INTO user_daily_counters (daily_count, last_updated, user_id) VALUES(?,?,?) 
		 ON DUPLICATE KEY UPDATE daily_count = ?, last_updated = ?
	`
	_, err := s.db.Exec(query, user.DailyCounter.DailyCount, user.DailyCounter.LastUpdated, user.ID,
		user.DailyCounter.DailyCount, user.DailyCounter.LastUpdated)
	if err != nil {
		log.Printf("Failed to update user with error: '%v'.", err)
		return nil, err
	}

	return user, nil
}
