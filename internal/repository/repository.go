// Package repository
// Experiment repository to use go's new features Typed parameters
// Use as Database access layer for all models/entities
package repository

import "gorm.io/gorm"

// Create records
func Create[T comparable](db *gorm.DB, record T) (T, error) {
	var data T
	if err := db.Save(&record).Error; err != nil {
		return data, err
	}
	return record, nil
}

// GetByID get record by id
func GetByID[T comparable, IDType comparable](db *gorm.DB, id IDType) (T, error) {
	var data T
	if err := db.First(&data, id).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetAll records
func GetAll[T comparable](db *gorm.DB) ([]T, error) {
	var arrData []T
	db.Find(&arrData)
	return arrData, nil
}

// GetByIDs by given ids
func GetByIDs[T comparable, IDType comparable](db *gorm.DB, ids IDType) ([]T, error) {
	var records []T
	if err := db.Find(&records, ids).Error; err != nil {
		return records, err
	}
	return records, nil
}

// Update record
func Update[T comparable](db *gorm.DB, record T) (T, error) {
	db.First(&record)
	db.Save(&record)
	return record, nil
}

// Delete records
func Delete[T comparable](db *gorm.DB, record T) (T, error) {
	db.Delete(&record)
	return record, nil
}


