package database

import "gorm.io/gorm"

type Responstory interface {
	Insert(data interface{}) error
	// Inserts(datas []interface{}) error
	// Delete(id string) error
	Get(data interface{}, id string) error
	GetAll(data interface{}) error
}

type responstory struct {
	table *gorm.DB
}

func InitResponstory(client *gorm.DB, tablename string) Responstory {
	table := client.Table(tablename)
	newresponstory := responstory{table: table}
	return newresponstory
}

func (res responstory) Insert(data interface{}) error {
	if result := res.table.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

func (res responstory) Get(data interface{}, id string) error {
	result := res.table.Where("id = ?", id).Scan(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (res responstory) GetAll(data interface{}) error {
	result := res.table.Where("is_active", true).Scan(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
