package postgres

import (
	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
)

type PGTaskRepository struct {
	PGRepository
}

func NewPGTaskRepository(dbConn *sqlx.DB) *PGTaskRepository {
	return &PGTaskRepository{PGRepository{dbConn}}
}

func (t *PGTaskRepository) CreateTaskForUser(userID int, taskParam d.TaskCreateParam) (*d.Task, error) {
	task := d.Task{UserID: userID, Content: taskParam.Content}
	_, err := t.DBConn.NamedExec(
		"INSERT INTO tasks (user_id, content) values (:user_id, :content)",
		task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *PGTaskRepository) GetTasksForUser(userID int, dateStr string) ([]*d.Task, error) {
	rows, err := t.DBConn.Queryx(
		"SELECT * FROM tasks WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3",
		userID, dateStr+" 00:00:00", dateStr+" 23:59:59")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tasks := []*d.Task{}
	for rows.Next() {
		t := d.Task{}
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (t *PGTaskRepository) GetTaskCount(userID int, dateStr string) (int, error) {
	var count int
	err := t.DBConn.Get(&count,
		"SELECT count(id) FROM tasks WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3",
		userID, dateStr+" 00:00:00", dateStr+" 23:59:59")
	if err != nil {
		return 0, err
	}

	return count, nil
}
