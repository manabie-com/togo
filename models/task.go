package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/khoale193/togo/models/dbcon"
	"github.com/khoale193/togo/pkg/e"
)

type Task struct {
	ID        int64         `gorm:"column:id;primary_key;auto_increment" db:"id"`
	Name      string        `gorm:"column:name" db:"name"`
	MemberID  int           `gorm:"column:member_id" db:"member_id"`
	CreatedAt time.Time     `gorm:"column:created_at" db:"created_at"`
	CreatedBy sql.NullInt64 `gorm:"column:created_by;default:null" db:"created_by"`
	UpdatedAt sql.NullTime  `gorm:"column:updated_at;default:null" db:"updated_at"`
	UpdatedBy sql.NullInt64 `gorm:"column:updated_by;default:null" db:"updated_by"`
	DeletedAt sql.NullTime  `gorm:"column:deleted_at;default:null" db:"deleted_at"`
	DeletedBy sql.NullInt64 `gorm:"column:deleted_by;default:null" db:"deleted_by"`
}

func (Task) TableName() string {
	return e.TaskTable
}

func (a *Task) CreateTask() (int64, error) {
	model := []Task{
		{Name: a.Name, MemberID: a.MemberID, CreatedAt: time.Now()},
	}
	columnName := []string{"Name", "MemberID", "CreatedAt"}
	if result, err := dbcon.GetSqlXDB().NamedExec(fmt.Sprintf("insert into %s (%s) values (%s)",
		(Task{}).TableName(),
		ColumnsName((Task{}).TableName(), &Task{}, columnName),
		ColumnsNameValueList(&Task{}, columnName),
	), model); err != nil {
		return 0, err
	} else if lastInsertId, err := result.LastInsertId(); lastInsertId != 0 && err == nil {
		return lastInsertId, nil
	} else {
		return 0, err
	}
}

func DeleteTaskByMemberID(memberID int) error {
	updateSQLInit := `delete from %[1]s where %[1]s.member_id = ?;`
	updateSQL := fmt.Sprintf(updateSQLInit, (Task{}).TableName())
	if result, err := dbcon.GetSqlXDB().Exec(updateSQL, memberID); err != nil {
		return err
	} else if affectedRows, err := result.RowsAffected(); err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return err
	} else {
		fmt.Printf("update data success, affected rows:%d\n", affectedRows)
	}
	return nil
}

func (Task) GetLatestInserted() (*Task, error) {
	selectSQLInit := `select * from %[1]s order by %[1]s.id desc limit 1;`
	selectSQL := fmt.Sprintf(selectSQLInit, (Task{}).TableName())
	var data Task
	if err := dbcon.GetSqlXDB().Get(&data, selectSQL); err != nil {
		return nil, err
	}
	return &data, nil
}
