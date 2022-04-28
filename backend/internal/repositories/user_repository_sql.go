package repositories

import (
	"database/sql"
	"context"
	"manabie.com/internal/models"
)

type UserRepositorySql struct {
	tx             *sql.Tx
}

func MakeUserRepositorySql(
	iTx *sql.Tx,
) UserRepositorySql {
	return UserRepositorySql{
		tx: iTx,
	}
}

func (r UserRepositorySql) FetchUserById(iContext context.Context, iId int) (models.User, error) {
	result, err := r.tx.QueryContext(iContext, `
		SELECT id, name, task_limit FROM "user" WHERE id=$1
	`, iId)

	if err != nil {
		return models.User{}, err
	}
	defer result.Close()

	result.Next()
	var id int
	var name string
	var maxNumberOfTasks int
	result.Scan(
		&id,
		&name,
		&maxNumberOfTasks,
	)

	return models.MakeUser(id, name, maxNumberOfTasks,), nil

}