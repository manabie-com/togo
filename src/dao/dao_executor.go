package dao

import (
	"github.com/HoangMV/togo/lib/pgsql"
	"github.com/HoangMV/togo/src/models/entity"
)

func (dao *DAO) InsertUser(obj *entity.User) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	INSERT INTO 
		users(username, password) 
	VALUES($1, $2)
	RETURNING id;`

	return dao.db.QueryRowContext(ctx, query, obj.Username, obj.Password).Scan(&obj.ID)
}

func (dao *DAO) InsertTodo(obj *entity.Todo) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	INSERT INTO 
		todos(user_id, content, status) 
	VALUES
		($1, $2, $3)
	RETURNING id;`

	return dao.db.QueryRowContext(ctx, query, obj.UserID, obj.Content, obj.Status).Scan(&obj.ID)
}

func (dao *DAO) InsertUserMaxTodo(obj *entity.UserTodoConfig) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	INSERT INTO 
		user_todo_config(user_id, max_todo) 
	VALUES
		($1, $2)
	RETURNING user_id;`

	return dao.db.QueryRowContext(ctx, query, obj.UserID, obj.MaxTodo).Err()
}

func (dao *DAO) UpdateTodo(obj *entity.Todo) error {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	UPDATE 
		todos 
	SET
		content = $1, status =$2
	WHERE
		id=$3;`

	_, err := dao.db.ExecContext(ctx, query, obj.UserID, obj.Content, obj.Status)
	return err
}

func (dao *DAO) GetUserByUsername(username string) (*entity.User, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	SELECT 
		id, username, password, created_at, updated_at
	FROM 
		users
	WHERE
		username = $1; 
	`
	user := &entity.User{}
	err := dao.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (dao *DAO) SelectTodoByUserID(userID int64) (*entity.Todo, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	SELECT
		id, user_id, content, status, created_at, updated_at
	FROM
		todos
	WHERE
		user_id = $1;
	`

	todo := &entity.Todo{}
	err := dao.db.SelectContext(ctx, &todo, query, userID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (dao *DAO) CountUserTodoInCurrentDay(userID int64) (int, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	SELECT 
		COUNT(*)
	FROM 
		todos
	WHERE
		user_id = $1 
		AND created_at <  date_trunc('day', now()) + interval '1 day'
		AND created_ad >= date_trunc('day', now()); 
	`
	var count int
	err := dao.db.GetContext(ctx, &count, query, userID)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (dao *DAO) GetMaxUserTodoOneDay(userID int64) (int, error) {
	ctx, cancel := pgsql.GetDefaultContext()
	defer cancel()

	query := `
	SELECT 
		max_todo
	FROM 
		user_todo_config
	WHERE
		user_id = $1; 
	`
	var count int
	err := dao.db.GetContext(ctx, &count, query, userID)
	if err != nil {
		return -1, err
	}

	return count, nil
}
