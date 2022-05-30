package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
	"togo/connections"
)

type Task struct {
	Id           	int64           	`json:"id" gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	Summary         string          	`json:"summary" gorm:"column:summary;NOT NULL"`
	Description     string          	`json:"description" gorm:"column:description;NULL"`
	Assignee     	int64          	`json:"assignee" gorm:"column:assignee;NOT NULL"`
	TaskDate     	string        	`json:"taskDate" gorm:"column:task_date;NULL"`
	CreatedAt     	time.Time        	`json:"-" gorm:"<-:false;column:created_at;NULL"`
}

type DailyLimit struct {
	Id           	int64           	`json:"id" gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	UserId         	int64          	`json:"UserId" gorm:"column:user_id;NOT NULL"`
	TaskDate     	string        	`json:"taskDate" gorm:"column:task_date;NOT NULL"`
	TaskLimit       int          		`json:"taskLimit" gorm:"column:task_limit;NOT NULL"`
	CreatedAt     	time.Time        	`json:"-" gorm:"<-:false;column:created_at;NULL"`
	UpdatedBy     	string         		`json:"-" gorm:"column:updated_by;NULL"`
}

func (Task) TableName() string {
	return "tasks"
}

func (DailyLimit) TableName() string {
	return "daily_limit"
}

var db *gorm.DB
var sqlDB *sql.DB

func init() {
	connect()
}

func connect() {
	var err error
	if db, err = connections.Connect(); err != nil {
		panic(err)
	}
	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
}

func (t *Task) Submit() (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		connect()
	}
	var taskId int64
	tx, _ := db.DB()
	/* Function submit_task
		Input: summary TEXT, description TEXT, assignee BIGINT, task_ate DATE)
		Output: id BIGINT
		Using pg_advisory_lock(userId) to block the possibility of 2 tasks of a user being added at the same time
	*/
	err = tx.QueryRow("SELECT submit_task($1,$2,$3,$4);", t.Summary, t.Description, t.Assignee, t.TaskDate).Scan(&taskId)
	if err == nil && taskId == 0 {
		err = errors.New("limit tasks exceeded")
	}
	return
}

func (t *Task) IsValidTask() bool {
	if t.Summary == "" || t.Assignee == 0 || t.TaskDate == "" {
		return false
	}
	if _, err := time.Parse("2006-01-02 15:04", t.TaskDate + " 15:04"); err != nil {
		return false
	}
	return true
}