package store

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type fileStore struct {
	Mu   sync.Mutex
	Data [][]string
}

var FileStoreConfig struct {
	DataFilePath string
	Fs           fileStore
}

type Store interface {
	InsertTask(id, userID, name, createdAt string) error
	GetTaskByID(id string) ([4]string, error)
}

func NewStore() (Store, error) {
	dataFilePath := "./data.csv"
	_, err := os.Stat(dataFilePath)

	if err != nil {
		_, err := os.Create(dataFilePath)
		if err != nil {
			return nil, err
		}
	}
	FileStoreConfig.Fs = fileStore{Mu: sync.Mutex{}, Data: make([][]string, 0)}
	FileStoreConfig.DataFilePath = dataFilePath
	return &FileStoreConfig.Fs, nil
}

func (j *fileStore) InsertTask(id, userID, name, createdAt string) error {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	return j.writeToFile([4]string{id, userID, name, createdAt})
}

func (j *fileStore) GetTaskByID(id string) ([4]string, error) {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	err := j.readFromFile()
	if err != nil {
		return [4]string{}, err
	}

	for _, task := range j.Data {
		if task[0] == id {
			return [4]string{task[0], task[1], task[2], task[3]}, nil
		}
	}

	return [4]string{}, nil
}

func (j *fileStore) writeToFile(data [4]string) error {
	f, err := os.OpenFile(FileStoreConfig.DataFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	// defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", data[0], data[1], data[2], data[3]))
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (j *fileStore) readFromFile() error {
	f, err := os.Open(FileStoreConfig.DataFilePath)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(f)

	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	j.Data = data

	return nil
}
