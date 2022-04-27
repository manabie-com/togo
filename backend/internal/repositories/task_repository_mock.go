package repositories

import (
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"context"
)

type TaskRepositoryMock struct {
	count int
	Tasks map[int][]models.Task
	userRepository *UserRepositoryMock
}

func MakeTaskRepositoryMock(iUserRepoitory *UserRepositoryMock) TaskRepositoryMock {
	return TaskRepositoryMock {
		count: 0,
		Tasks: map[int][]models.Task{},
		userRepository: iUserRepoitory,
	}
}

func (r *TaskRepositoryMock) CreateTaskForUser(
	iContext context.Context, 
	iUser models.User, 
	iTasks []models.Task,
) ([]models.Task, error) {
	if _, ok := r.userRepository.Users[iUser.Id]; !ok {
		return []models.Task{}, common.NotFound
	}

	ret := []models.Task{}
	for _, task := range iTasks {
		task.Id = r.count
		task.Owner = &iUser
		ret = append(ret, task)
		r.count += 1
	}
	r.Tasks[iUser.Id] = append(r.Tasks[iUser.Id], ret...)
	return ret, nil
}

func (r *TaskRepositoryMock) FetchNumberOfTaskForUser(
	iContext context.Context, 
	iUser models.User,
) (int, error) {
	if _, ok := r.userRepository.Users[iUser.Id]; !ok {
		return 0, common.NotFound
	}

	return len(r.Tasks[iUser.Id]), nil
}

func (r *TaskRepositoryMock) FetchNumberOfTaskForUserCreatedOnDay(
	iContext context.Context, 
	iUser models.User, 
	iCreatedTime common.Time,
) (int, error) {
	if _, ok := r.userRepository.Users[iUser.Id]; !ok {
		return 0, common.NotFound
	}

	searchDay := iCreatedTime.Year() * 10000 + int(iCreatedTime.Month()) * 100 + iCreatedTime.Day()
	count := 0
	for _, task := range r.Tasks[iUser.Id] {
		day := task.CreatedTime.Year() * 10000 + int(task.CreatedTime.Month()) * 100 + task.CreatedTime.Day()
		if day == searchDay {
			count += 1
		}
	}

	return count, nil
}
