package services_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/services"
)

type MockedTaskRepo struct {
	Tasks []*entities.Task
}

func (m *MockedTaskRepo) CountTasksOfUserByDate(ctx context.Context, userID string, createdDate string) (int, error) {
	tasks, err := m.GetTasksByUserIDAndDate(context.TODO(), userID, createdDate)
	return len(tasks), err
}

func (m *MockedTaskRepo) GetTasksByUserIDAndDate(ctx context.Context, userID, createdDate string) ([]*entities.Task, error) {
	var result []*entities.Task
	for _, task := range m.Tasks {
		if task.CreatedDate == createdDate && task.UserID == userID {
			result = append(result, task)
		}
	}
	return result, nil
}

func (m *MockedTaskRepo) SaveTask(ctx context.Context, task entities.Task) (*entities.Task, error) {
	return &task, nil
}

func newDefaultTaskSvc(tasks []*entities.Task) *services.TaskSvc {
	return services.NewTaskService(services.TaskServiceConfiguration{
		TaskRepo: &MockedTaskRepo{Tasks: tasks},
	})
}

func assertTask(t testing.TB, taskGot, taskWant *entities.Task) {
	t.Helper()
	if match := reflect.DeepEqual(taskGot, taskWant); !match {
		t.Errorf("Got user is different from user that exist, got: %v, want: %v", taskGot, taskWant)
	}
}

func TestTaskSvc(t *testing.T) {
	defaultTasks := []*entities.Task{{ID: "a18a52c0-e2ea-4e66-97fb-531b1936c72b", Content: "Hello world", UserID: "phuonghau"}}

	t.Run("AddTask should performs saving action against database", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		newTask := entities.Task{ID: "", Content: "Foo bar", UserID: "phuonghau"}
		savedTask, err := taskSvc.AddTask(context.TODO(), newTask)
		assertError(t, err, nil)
		if savedTask != nil {
			// Get random id of saved task
			newTask.ID = savedTask.ID

			// Get generated date
			newTask.CreatedDate = savedTask.CreatedDate
			assertTask(t, savedTask, &newTask)
		}
	})

	t.Run("AddTask should have uuid as default id", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		newTask := entities.Task{ID: "pseudoID", Content: "Foo bar", UserID: "phuonghau"}

		savedTask, err := taskSvc.AddTask(context.TODO(), newTask)
		assertError(t, err, nil)
		if savedTask.ID == newTask.ID {
			t.Errorf("Task saved but Task.Id was not generated, still got: %v", savedTask.ID)
		}
	})

	t.Run("AddTask save task with have field created_date in format of yyyy-dd-mm", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		newTask := entities.Task{ID: "pseudoID", Content: "Foo bar", UserID: "phuonghau"}

		savedTask, err := taskSvc.AddTask(context.TODO(), newTask)
		assertError(t, err, nil)
		if len(savedTask.CreatedDate) == 0 {
			t.Errorf("Created task doesn't contain created_date")
		}
	})
}

func TestListTasksByUserAndDate(t *testing.T) {
	tasksOnMarch112021 := []*entities.Task{
		{ID: "a18a52c0-e2ea-4e66-97fb-531b1936c72b", Content: "C1", UserID: "phuonghau", CreatedDate: "2021-03-11"},
		{ID: "a18a52c0-e2ea-4e66-97fb-531b1936c36k", Content: "C2", UserID: "phuonghau", CreatedDate: "2021-03-11"},
	}
	defaultTasks := []*entities.Task{
		{ID: "a18a52c0-e2ea-4e66-97fb-531b1936c24e", Content: "C3", UserID: "phuonghau", CreatedDate: "2021-03-09"},
	}

	defaultTasks = append(defaultTasks, tasksOnMarch112021...)

	t.Run("Should return all tasks on specific date of current store", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		tasks, err := taskSvc.ListTasksByUserAndDate(context.TODO(), "phuonghau", "2021-03-11")
		assertError(t, err, nil)
		if !reflect.DeepEqual(tasksOnMarch112021, tasks) {
			t.Errorf("Tasks was not retrieved as expected, got: %v, want: %v", tasks, tasksOnMarch112021)
		}
	})

	t.Run("Should return error when provide invalid date", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		invalidDates := []string{"-3220923", "2020-30-30", "2020-15-15", "2232-3232-32323"}
		for _, inValidDate := range invalidDates {
			tasks, err := taskSvc.ListTasksByUserAndDate(context.TODO(), "phuonghau", inValidDate)
			if len(tasks) != 0 {
				t.Errorf("Provide error date, but still go: %v", tasks)
			}
			assertError(t, err, entities.ErrTaskInvalidCreatedDate)
		}
	})

	t.Run("Should return null if no task found on provided date", func(t *testing.T) {
		taskSvc := newDefaultTaskSvc(defaultTasks)
		dateThatHaveNoTask := "2022-03-03"
		tasks, err := taskSvc.ListTasksByUserAndDate(context.TODO(), "phuonghau", dateThatHaveNoTask)
		if len(tasks) != 0 {
			t.Errorf("Should return no task on %s, but still got: %v", dateThatHaveNoTask, tasks)
		}
		assertError(t, err, nil)
	})
}
