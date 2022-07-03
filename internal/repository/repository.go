// Package repository
// Experiment repository to use go's new features Typed parameters
// Use as Database access layer for all models/entities
package repository

import "github.com/xrexonx/togo/cmd/app/config/database"

// Create records
func Create[T comparable](record T) (T, error) {
	var data T
	if err := database.Instance.Save(&record).Error; err != nil {
		return data, err
	}
	return record, nil
}

// GetByID get record by id
func GetByID[T comparable, IDType comparable](id IDType) (T, error) {
	var data T
	if err := database.Instance.First(&data, id).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetAll records
func GetAll[T comparable]() ([]T, error) {
	var arrData []T
	database.Instance.Find(&arrData)
	return arrData, nil
}

// GetByIDs by given ids
func GetByIDs[T comparable, IDType comparable](ids IDType) ([]T, error) {
	var records []T
	if err := database.Instance.Find(&records, ids).Error; err != nil {
		return records, err
	}
	return records, nil
}

// Update record
func Update[T comparable](record T) (T, error) {
	database.Instance.First(&record)
	database.Instance.Save(&record)
	return record, nil
}

// Delete records
func Delete[T comparable](record T) (T, error) {
	database.Instance.Delete(&record)
	return record, nil
}
