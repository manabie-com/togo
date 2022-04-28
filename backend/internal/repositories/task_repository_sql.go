package repositories

import (
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type TaskRepositorySql struct {
	tx *sql.Tx
}

func MakeTaskRepositorySql(
	iTx *sql.Tx,
) TaskRepositorySql {
	return TaskRepositorySql {
		tx: iTx,
	}
}

func (r TaskRepositorySql) CreateTaskForUser(iContext context.Context, iUser models.User, iTasks []models.Task) ([]models.Task, error) {
	argsString := []string{}
	args := []interface{}{}
	count := 1

	for _, task := range iTasks {
		argsString = append(argsString, fmt.Sprintf("($%d, $%d, $%d, $%d)", count, count + 1, count + 2, count + 3))
		args = append(args, iUser.Id, task.Title, task.Content, task.CreatedTime)
		count += 4
	}

	statement := `INSERT INTO "tasks" (
			user_id,
			title,
			content,
			created_time
		) VALUES
	`
	statement += strings.Join(argsString, ",")
	statement += " RETURNING id"

	result, err := r.tx.QueryContext(iContext, statement, args...)
	if err != nil {
		return []models.Task{}, err
	}
	defer result.Close()

	ret := make([]models.Task, len(iTasks))
	copy(ret, iTasks)
	index := 0
	for result.Next() {
		var id int
		result.Scan(&id)
		ret[index].Id = id
		ret[index].Owner = &iUser
		index += 1
	}


	return ret, nil
}

func (r TaskRepositorySql)FetchNumberOfTaskForUser(iContext context.Context, iUser models.User) (int, error) {
	result, err := r.tx.QueryContext(iContext, `
		SELECT COUNT(*) FROM "tasks" WHERE user_id=$1
	`, iUser.Id)

	if err != nil {
		return 0, err
	}
	defer result.Close()

	result.Next()
	var count int
	result.Scan(&count)
	return count, nil
}


func (r TaskRepositorySql)FetchTasksForUser(iContext context.Context, iUser models.User) ([]models.Task, error) {
	ret := []models.Task{}
	result, err := r.tx.QueryContext(iContext, `
		SELECT id, title, content, created_time FROM "tasks" WHERE user_id=$1 ORDER BY created_time DESC, id DESC
	`, iUser.Id)
	if err != nil {
		return []models.Task{}, err
	}
	defer result.Close()

	for result.Next() {
		var id int
		var title, content string
		var createdTime common.Time
		result.Scan(
			&id,
			&title,
			&content,
			&createdTime,
		)

		task := models.MakeTask(
			id,
			title,
			content,
			createdTime,
			&iUser,
		)

		ret = append(ret, task)
	}
	return ret, nil
}

func (r TaskRepositorySql)FetchNumberOfTaskForUserCreatedOnDay(iContext context.Context, iUser models.User, iCreatedTime common.Time) (int, error) {
    year, month, day := iCreatedTime.Date()
	startOfCurrentDay := time.Date(year, month, day, 0, 0, 0, 0, iCreatedTime.Location())
	startOfNextDay := time.Date(year, month, day + 1, 0, 0, 0, 0, iCreatedTime.Location())
	result, err := r.tx.QueryContext(iContext, `
		SELECT COUNT(*) FROM "tasks" WHERE user_id=$1 AND created_time >= $2 AND created_time < $3 
	`, iUser.Id, startOfCurrentDay, startOfNextDay)

	if err != nil {
		return 0, err
	}
	defer result.Close()

	result.Next()
	var count int
	result.Scan(&count)
	return count, nil
}
