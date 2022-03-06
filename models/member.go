package models

import (
	"database/sql"
	"fmt"
	"time"
	"togo/models/dbcon"

	"togo/pkg/e"
)

type Member struct {
	ID        int64         `gorm:"column:id;primary_key;auto_increment" db:"id"`
	Username  string        `gorm:"column:nick_name" db:"nick_name"`
	Password  string        `gorm:"column:password" db:"password"`
	MaxTask   int           `gorm:"column:max_task;default:null" db:"max_task"`
	Status    bool          `gorm:"column:status" db:"status"`
	CreatedAt time.Time     `gorm:"column:created_at" db:"created_at"`
	CreatedBy sql.NullInt64 `gorm:"column:created_by;default:null" db:"created_by"`
	UpdatedAt sql.NullTime  `gorm:"column:updated_at;default:null" db:"updated_at"`
	UpdatedBy sql.NullInt64 `gorm:"column:updated_by;default:null" db:"updated_by"`
	DeletedAt sql.NullTime  `gorm:"column:deleted_at;default:null" db:"deleted_at"`
	DeletedBy sql.NullInt64 `gorm:"column:deleted_by;default:null" db:"deleted_by"`
}

func (Member) TableName() string {
	return e.MemberTable
}

func (a Member) FindByUsername(username string) (*Member, error) {
	sqlXDB := dbcon.GetSqlXDB()
	selectSQLInit := "select * from %[1]s where binary %[1]s.%[2]s = binary ?;"
	selectSQL := fmt.Sprintf(selectSQLInit, (Member{}).TableName(), ColName(&Member{}, "Username"))
	var model Member
	if err := sqlXDB.Get(&model, selectSQL, username); err != nil {
		return nil, err
	}
	return &model, nil
}
