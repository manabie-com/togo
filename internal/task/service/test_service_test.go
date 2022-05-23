package service

import (
	"testing"
)

func TestCreate(t *testing.T) {
	//mockTaskRepo := new(mocks.ITaskRepository)
	//mockUserRepo := new(mocks.IUserRepository)
	//taskParams := domain.TaskParams{
	//	Content:   "This is a task",
	//	UserEmail: "fake-email" + strconv.Itoa(rand.Intn(1000)) + "@gmail.com",
	//}
	//
	//t.Run("success", func(t *testing.T) {
	//	tempMockParams := taskParams
	//	mockTaskRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil).Once()
	//
	//	taskService := NewTaskService(mockTaskRepo, mockUserRepo)
	//	createdTask, err := taskService.Create(tempMockParams)
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, createdTask.Content, tempMockParams.Content)
	//	mockTaskRepo.AssertExpectations(t)
	//})
}
