package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Responstory interface {
	Insert(data interface{}) error
	// Inserts(datas []interface{}) error
	// Delete(id string) error
	Get(data interface{}, id string) error
	Find(data interface{}, query string, args ...interface{}) error
	GetAll(data interface{}) error
}

type responstory struct {
	client *gorm.DB
	table  string
}

func InitResponstory(client *gorm.DB, tablename string) Responstory {
	newresponstory := responstory{client: client, table: tablename}
	return newresponstory
}

func (res responstory) Insert(data interface{}) error {
	if result := res.client.Table(res.table).Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

func (res responstory) Get(data interface{}, id string) error {
	fmt.Println(res.client.Table(res.table).Statement.Vars...)
	result := res.client.Table(res.table).Where("id = ?", id).Scan(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (res responstory) Find(data interface{}, query string, args ...interface{}) error {
	result := res.client.Table(res.table).Where(query, args...).Scan(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (res responstory) GetAll(data interface{}) error {
	result := res.client.Table(res.table).Where(res.table).Where("is_active", true).Scan(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
