package sdkcm

import (
	"database/sql/driver"
	"errors"
	"time"
)

// For reading
type SQLModel struct {
	ID        int        `json:"id" gorm:"id,PRIMARY_KEY"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (m *SQLModel) FullFill() {
	t := time.Now()

	if m.UpdatedAt == nil {
		m.UpdatedAt = &t
	}
}

func NewSQLModel() *SQLModel {
	t := time.Now()
	return &SQLModel{
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

func NewUpsertSQLModel(id int) *SQLModel {
	t := time.Now()

	return &SQLModel{
		ID:        id,
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

func NewUpsertWithoutIdSQLModel() *SQLModel {
	t := time.Now()

	return &SQLModel{
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

// Set time format layout. Default: 2006-01-02
func SetDateFormat(layout string) {
	dateFmt = layout
}

type JSON []byte

// This method for mapping JSON to json data type in sql
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// This method for scanning JSON from json data type in sql
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	if s, ok := value.([]byte); ok {
		*j = append((*j)[0:0], s...)
		return nil
	}

	return errors.New("invalid Scan Source")
}

func (j *JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return nil, errors.New("object json is nil")
	}

	return *j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("object json is nil")
	}

	*j = JSON(data)
	return nil
}
