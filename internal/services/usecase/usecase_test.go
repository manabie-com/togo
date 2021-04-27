package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/services/auth"
	"github.com/manabie-com/togo/internal/services/usecase"
	"github.com/manabie-com/togo/internal/storages/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthToken(t *testing.T) {
	db := models.Connect()
	handler := usecase.NewUsecase(db)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Store = mock.NewMockStorageInterface(ctrl)

	t.Run("Success", func(t *testing.T) {
		handler.Store.(*mock.MockStorageInterface).EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.User{
			ID:       100,
			Username: "test123",
			Password: "pwd123",
		}, nil)
		//create token with user ID
		token, _ := auth.CreateToken(100)
		// call the usecase
		expected, err := handler.GetAuthToken(context.Background(), "test123", "pwd123")
		//testing
		assert.Equal(t, err, nil, "they should be equal")
		assert.Equal(t, expected, token, "they should be equal")
	})

	t.Run("Incorrect usename or password", func(t *testing.T) {
		handler.Store.(*mock.MockStorageInterface).EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.User{}, errors.New("some errors"))
		// call the usecase
		_, err := handler.GetAuthToken(context.Background(), "test123", "pwd123")
		//testing
		assert.Equal(t, err, errors.New("incorrect username or password"), "they should be equal")
	})

	t.Run("Create token error", func(t *testing.T) {
		handler.Store.(*mock.MockStorageInterface).EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.User{
			ID:       100,
			Username: "test123",
			Password: "pwd123",
		}, nil)
		//create token with user ID

		createToken := func(id int) (string, error) {
			return "", errors.New("there was an error generating the API token")
		}
		token, tkErr := createToken(100)
		// call the usecase
		expected, _ := handler.GetAuthToken(context.Background(), "", "")
		//testing
		assert.EqualError(t, tkErr, "there was an error generating the API token")
		assert.NotEqual(t, expected, token, "they should be equal")

	})

	t.Run("Create token error with other UserID", func(t *testing.T) {
		handler.Store.(*mock.MockStorageInterface).EXPECT().ValidateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.User{
			ID:       100,
			Username: "test123",
			Password: "pwd123",
		}, nil)
		//create token with user ID

		token, _ := auth.CreateToken(200)

		// call the usecase
		expected, _ := handler.GetAuthToken(context.Background(), "test123", "pwd123")
		//testing
		assert.NotEqual(t, expected, token, "they should be equal")
	})
}

func TestRetrieveTasks(t *testing.T) {
	t.Parallel()
	db := models.Connect()
	handler := usecase.NewUsecase(db)

	// create the task mock interface
	ctrl := gomock.NewController(t)
	handler.Store = mock.NewMockStorageInterface(ctrl)
	t.Run("Success", func(t *testing.T) {
		allTasks := make([]*models.Task, 0)
		allTasks = append(allTasks, &models.Task{
			ID:          1,
			Content:     "statement 1",
			UserID:      1,
			CreatedDate: "2021-04-26",
		})
		allTasks = append(allTasks, &models.Task{
			ID:          2,
			Content:     "statement 2",
			UserID:      2,
			CreatedDate: "2021-04-26",
		})
		//do storage method
		handler.Store.(*mock.MockStorageInterface).EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(allTasks, nil)
		// call the usecase method
		ctx := context.Background()
		tasks, error := handler.RetrieveTasks(ctx, 999, "2021-04-26")
		//testing
		assert.NoError(t, error)
		assert.Equal(t, tasks, allTasks, "they should be equal")
	})
	t.Run("Failed", func(t *testing.T) {
		allTasks := make([]*models.Task, 0)
		//do storage method
		handler.Store.(*mock.MockStorageInterface).EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(allTasks, errors.New("some errors"))
		// call the usecase method
		ctx := context.Background()
		tasks, error := handler.RetrieveTasks(ctx, 999, "2021-04-26")
		//testing
		assert.Equal(t, error, errors.New("failure to retrieve tasks"), "they should be equal")
		assert.NotEqual(t, tasks, allTasks, "they should be equal")
	})
}
func TestAddTask(t *testing.T) {
	db := models.Connect()
	ctx := context.Background()
	handler := usecase.NewUsecase(db)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Store = mock.NewMockStorageInterface(ctrl)

	t.Run("Sucess", func(t *testing.T) {
		// Data
		taskAct := &models.Task{
			ID:          100,
			Content:     "Statement 1",
			UserID:      200,
			CreatedDate: time.Now().Format("2006-01-02"),
		}
		//do storage method
		handler.Store.(*mock.MockStorageInterface).EXPECT().AddTask(gomock.Any(), taskAct).Return(nil)
		// call the usecase method
		taskExpect, error := handler.AddTask(ctx, 200, &models.Task{
			ID:          100,
			Content:     "Statement 1",
			UserID:      200,
			CreatedDate: time.Now().Format("2006-01-02"),
		})
		//included to facilitate the check with the returned value
		taskAct.ID = 100
		//testing
		assert.NoError(t, error)
		assert.Equal(t, taskExpect, taskAct, "they should be equal")
	})

	t.Run("Failed", func(t *testing.T) {
		// Data
		taskAct := &models.Task{
			Content: "Statement 1",
			UserID:  1,
		}
		//do the storage method
		handler.Store.(*mock.MockStorageInterface).EXPECT().AddTask(gomock.Any(), taskAct).Return(errors.New("the task daily limit is reached"))
		// call the usecase service
		task, error := handler.AddTask(ctx, 200, taskAct)
		assert.Equal(t, error, errors.New("the task daily limit is reached"), "they should be equal")
		assert.Nil(t, task)
	})
}
