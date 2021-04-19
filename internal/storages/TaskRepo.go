package storages

import (
	"fmt"
)

//TASK//

type ITaskRepo interface {
	StoreFromRepo(task Task) error
	FindByIdAndTimeFromRepo(user_id, time string) ([]Task, error)	
}

type TaskRepo struct {
	IDBHandler
}

func (taskRepo *TaskRepo) FindByIdAndTimeFromRepo(user_id, time string) ([]Task, error) {

	rows, err := taskRepo.Query(fmt.Sprintf("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = '%s' AND created_date = '%s'", user_id, time))
	if err != nil {
		return nil, err
	}	
	
	var tasks []Task
	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	// if err := (); err != nil {
	// 	return nil, err
	// }

	return tasks, nil	
}

func (taskRepo *TaskRepo) StoreFromRepo(t Task) error{
	err := taskRepo.Execute(fmt.Sprintf("INSERT INTO tasks (id, content, user_id, created_date) VALUES ('%s', '%s', '%s', '%s')", t.ID, t.Content, t.UserID, t.CreatedDate))	
	if err != nil {
		return err
	}

	return nil
}
