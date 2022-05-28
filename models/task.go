package models

import (
	"database/sql"
	"errors"
	"github.com/beego/beego/v2/server/web/context"
	"gorm.io/gorm"
	"log"
	"time"
	"togo/connections"
)

type Task struct {
	Id           	int64           	`json:"id" gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	Summary         string          	`json:"summary" gorm:"column:summary;NOT NULL"`
	Description     string          	`json:"description" gorm:"column:description;NULL"`
	Assignee     	string          	`json:"assignee" gorm:"column:assignee;NOT NULL"`
	TaskDate     	string        	`json:"taskDate" gorm:"column:task_date;NULL"`
	CreatedAt     	time.Time        	`json:"createdAt,omitempty" gorm:"<-:false;column:created_at;NULL"`
}

type DailyLimit struct {
	Id           	int64           	`json:"id" gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	UID         	string          	`json:"uid" gorm:"column:uid;NOT NULL"`
	TaskDate     	string        	`json:"taskDate" gorm:"column:task_date;NOT NULL"`
	TaskLimit       int          		`json:"taskLimit" gorm:"column:task_limit;NOT NULL"`
	CreatedAt     	time.Time        	`json:"createdAt,omitempty" gorm:"<-:false;column:created_at;NULL"`
	UpdatedBy     	string         		`json:"updatedBy,omitempty" gorm:"column:updated_by;NULL"`
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

func (t *Task) Submit(ctx *context.Context) (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		connect()
	}
	var daily DailyLimit
	var tasks []Task
	// Todo: implement user management to get default number of daily tasks
	limit := 10
	// Todo: try to implement db function
	return db.Transaction(func(tx *gorm.DB) error {
		select {
		case <-ctx.Request.Context().Done():
			log.Println(ctx.Request.Context().Err())
			return ctx.Request.Context().Err()
		default:
			if err = tx.Where(&DailyLimit{UID: t.Assignee, TaskDate: t.TaskDate}).First(&daily).Error;
				!errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
				return err
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				limit = daily.TaskLimit
			}
			if tasks, err = countTasks(t.TaskDate, t.Assignee);  err != nil {
				return err
			}
			if len(tasks) >= limit {
				return errors.New("limit exceeded")
			}
			if err = tx.Create(t).Error; err != nil {
				return err
			}
			return nil
		}
	})
}

func (t *Task) IsValidTask() bool {
	if t.Summary == "" || t.Assignee == "" || t.TaskDate == "" {
		return false
	}
	if _, err := time.Parse("2006-01-02 15:04", t.TaskDate + " 15:04"); err != nil {
		return false
	}
	return true
}

func countTasks(taskDate string, uid string) (tasks []Task, err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		connect()
	}
	if err = db.Where(&Task{Assignee: uid, TaskDate: taskDate}).Find(&tasks).Error;  err != nil {
		return nil, err
	}
	return
}