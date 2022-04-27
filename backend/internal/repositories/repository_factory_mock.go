package repositories

import (
	"manabie.com/internal/common"
	"manabie.com/internal/models"
	"context"
	"fmt"
)

type RepositoryFactoryMock struct {
	Count int
	validIds map[int]bool
	TransactionLevelsHistory map[int]TransactionLevel
	TaskRepository *TaskRepositoryMock
	UserRepository *UserRepositoryMock
}

func (f *RepositoryFactoryMock) InitUsers(iNumberOfUsers int) {
	for i := 0; i < iNumberOfUsers; i++ {
		user := models.MakeUser(i, fmt.Sprintf("user-%d", i+1), i + 1)
		f.UserRepository.AddUser(user)
	}
}

func MakeRepositoryFactoryMock() *RepositoryFactoryMock {
	userRepository := MakeUserRepositoryMock()
	taskRepository := MakeTaskRepositoryMock(&userRepository)
	return &RepositoryFactoryMock {
		Count: 0,
		TransactionLevelsHistory: map[int]TransactionLevel{},
		validIds: map[int]bool{},
		TaskRepository: &taskRepository,
		UserRepository: &userRepository,
	}
}

func (m *RepositoryFactoryMock) StartTransactionAuto(
	iContext context.Context, 
	iIsolationLevel TransactionLevel,
	iHandler TransactionHandler,
) error {
	m.Count += 1
	m.validIds[m.Count] = true
	m.TransactionLevelsHistory[m.Count] = iIsolationLevel
	defer func () {
		delete(m.validIds, m.Count)
	}()
	return iHandler(TransactionId(m.Count))
}

func (m *RepositoryFactoryMock) GetTaskRepository(iId TransactionId) (TaskRepositoryI, error) {
	if _, ok := m.validIds[int(iId)]; !ok {
		return nil, common.NotFound
	}
	return m.TaskRepository, nil
}

func (m *RepositoryFactoryMock) GetUserRepository(iId TransactionId) (UserRepositoryI, error) {
	if _, ok := m.validIds[int(iId)]; !ok {
		return nil, common.NotFound
	}

	return m.UserRepository, nil
}

