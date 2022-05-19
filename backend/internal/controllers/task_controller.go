package controllers

import (
	"context"
	"manabie.com/internal/repositories"
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"errors"
)

type TaskController struct {
	factory repositories.RepositoryFactoryI
	clock common.ClockI
}

func MakeTaskController(
	iFactory repositories.RepositoryFactoryI,
	iClock common.ClockI,
) *TaskController {
	return &TaskController{
		factory: iFactory,
		clock: iClock,
	}
}

var TaskLimitExceeds = errors.New("TaskLimitExceeds")
/// returns true if exceeds
func (c TaskController) CreateTaskForUserId(
	iContext context.Context,
	iUserId int,
	iTaskTitle string,
	iTaskContent string,
) (models.Task, error) {
	var newTask models.Task
	/// need Serialzable isolation because otherwise the task limit may be exceeded
	/// we don't need performance anyway
	var err, txError error
	for true {
		err, txError = c.factory.StartTransactionAuto(
			iContext, 
			repositories.Serializable,
			func(iTransactionId repositories.TransactionId) error {
				userRepository, err := c.factory.GetUserRepository(iTransactionId)
				if err != nil {
					return err
				}

				taskRepository, err := c.factory.GetTaskRepository(iTransactionId)
				if err != nil {
					return err
				}

				user, err := userRepository.FetchUserById(iContext, iUserId)
				if err != nil {
					return err
				}

				currentTime := c.clock.Now()
				numberOfTasks, err := taskRepository.FetchNumberOfTaskForUserCreatedOnDay(iContext, user, currentTime)
				if err != nil {
					return err
				}

				if numberOfTasks >= user.MaxNumberOfTasks {
					return TaskLimitExceeds
				}

				task := models.MakeTask(
					-1,
					iTaskTitle,
					iTaskContent,
					currentTime,
					nil,
				)

				newTaskList, err := taskRepository.CreateTaskForUser(iContext, user, []models.Task{task})
				if err != nil {
					return err
				}

				newTask = newTaskList[0]
				return nil
			},
		)

		if txError != common.SqlSerializableTransactionError && err != common.SqlSerializableTransactionError {
			break
		}
	}

	if err != nil {
		return models.Task{}, err
	}

	if txError != nil {
		return models.Task{}, err
	}

	return newTask, nil
}
