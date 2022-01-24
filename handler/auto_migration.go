package handler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"strings"
)

type store struct {
	db *gorm.DB
	filePath string
}

func NewAutoMigration(filePath string, db *gorm.DB) *store {
	return &store{db: db, filePath: filePath}
}

func (s *store) Run (verion int) error {
	if err := s.configure(); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(s.filePath)
	if err != nil {
		return err
	}

	for _, f := range files {
		script, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", s.filePath, f.Name()))
		if err != nil {
			return err
		}

		multiScript := strings.Split(string(script), ";")
		for _, spt := range multiScript {
			spt = strings.TrimSuffix(spt, "\n")
			if spt == "" {
				continue
			}

			if err := s.db.Exec(spt).Error; err != nil {
				return err
			}

			log.Print(spt)
		}
	}

	return nil
}

func (s *store) configure () error {
	if s.db == nil {
		return errors.New("DatabaseIsNull")
	}

	if s.filePath == "" {
		return errors.New("FilePathIsNull")
	}

	return nil
}