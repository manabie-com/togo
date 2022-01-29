package database

import "gorm.io/gorm"

type Responstory interface {
	Insert(data interface{}) error
	// Inserts(datas []interface{}) error
	// Delete(id string) error
	// Get(id string) error
	// GetAll(id string) error
}

type responstory struct {
	client *gorm.DB
	table  interface{}
}

func InitResponstory(client *gorm.DB) Responstory {
	newresponstory := responstory{client: client}
	return newresponstory
}

func (res responstory) Insert(data interface{}) error {
	if result := res.client.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}
