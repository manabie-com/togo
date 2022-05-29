package persistent

import (
	"context"
	"database/sql"
	"log"
	"togo/domain/model"
	"togo/domain/repository"
)

type userMysqlRepo struct {
	db *sql.DB
}

func (this *userMysqlRepo) Create(ctx context.Context, u model.User) error {
	stmt, err := this.db.Prepare("insert into todo.tbl_user(username, password, limit_per_day) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}
	hashedPwd := u.Password
	_, err = stmt.ExecContext(ctx, u.Username, hashedPwd, u.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (this *userMysqlRepo) Get(ctx context.Context, username string) (u model.User, err error) {
	stmt, err := this.db.Prepare("select * from todo.tbl_user WHERE username = ? LIMIT 1")
	if err != nil {
		return u, err
	}
	rows, err := stmt.QueryContext(ctx, username)
	if err != nil {
		return u, err
	}

	defer rows.Close()
	for rows.Next() {
		var _id, _limit int
		var _username, _password string
		var _createdTime []uint8
		err = rows.Scan(&_id, &_username, &_password, &_limit, &_createdTime)
		if err != nil {
			return u, err
		}
		u = model.User{
			Username: _username,
			Id:       _id,
			Password: _password,
			Limit:    _limit,
		}
	}
	log.Printf("User: %#v", u)
	return u, nil
}

func NewUserMysqlRepository(db *sql.DB) repository.UserRepository {
	return &userMysqlRepo{
		db: db,
	}
}
