package persist

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"manabie.com/togo"
	"manabie.com/togo/utils"
	"time"
)

type storeImpl struct {
	db *sql.DB
}

type Config struct {
	HostAddr       string
	DB             string
	User           string
	Pass           string
	MaxOpenCnn     int
	MaxCnnLifeTime time.Duration
}

func NewStore(cfg Config) (togo.Store, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Pass, cfg.HostAddr, cfg.DB)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error %s when opening DB\n", err))
	}
	db.SetMaxOpenConns(cfg.MaxOpenCnn)
	db.SetConnMaxLifetime(cfg.MaxCnnLifeTime)
	if err = db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error %s when ping DB", err))
	}
	return &storeImpl{db: db}, nil
}

func (s *storeImpl) Initialize() error {
	var schema = `
		create table manabie_user (
			id int not null auto_increment primary key,
			username varchar(50) unique,
			password varchar(100),
			created_at timestamp not null,
			last_login timestamp
		);
		
		
		create table manabie_task (
			id int not null auto_increment primary key,
			task_name varchar(50) unique ,
			description varchar(500) ,
			created_at timestamp not null
		);
		
		create table manabie_task_limit_setting (
			user_id int not null,
			task_id int not null,
			num_limit int not null,
			updated_at timestamp not null ,
			constraint user_task_unique unique (user_id, task_id)
		);`

	_, err := s.db.Exec(schema)
	return err
}

func (s *storeImpl) CreateUser(username string, password string) (togo.UserEntity, error) {
	query := `insert into manabie_user (username, password, created_at) values (?, ?, ?)`

	now := utils.Now()
	res, err := s.db.Exec(query, username, password, now)
	if err != nil {
		return togo.UserEntity{}, errors.Wrapf(err, "fail to execute create user")
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return togo.UserEntity{}, errors.Wrapf(err, "fail to get row affected after create user")
	}

	return togo.UserEntity{
		int(lastInsertId), username, now, nil,
	}, nil
}

func (s *storeImpl) Login(username string, password string) (int, error) {
	query := `select id from manabie_user where username = ? and password = ?`
	tx, err := s.db.Begin()
	if err != nil {
		return 0, errors.Wrapf(err, "fail to begin transaction for login")
	}
	row := tx.QueryRow(query, username, password)

	var id int
	err = row.Scan(&id)
	if err != nil || id == 0 {
		return 0, errors.Wrapf(err, "username or password wrong for user %s", username)
	}

	updateLastLoginQuery := `update manabie_user set last_login = ? where username = ? and password = ?`
	if _, err = tx.Exec(updateLastLoginQuery, utils.Now(), username, password); err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			log.WithError(err).Errorf("can not rollback when update last_login for user %s", username)
		}
		return 0, errors.Wrapf(err, "fail to update last_login for user %s", username)
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Wrapf(err, "fail to commit login transactionfor user %s", username)
	}

	return id, nil
}

func (s *storeImpl) CreateTask(taskName string, description string) (togo.TaskEntity, error) {
	query := `insert into manabie_task (task_name, description, created_at) values (?, ?, ?)`
	now := utils.Now()
	res, err := s.db.Exec(query, taskName, description, now)
	if err != nil {
		return togo.TaskEntity{}, errors.Wrapf(err, "fail to execute create task")
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return togo.TaskEntity{}, errors.Wrapf(err, "fail to get last insert id after create task")
	}

	return togo.TaskEntity{int(lastInsertId), taskName, description, now}, nil
}

func (s *storeImpl) SetTaskLimit(userId int, taskId int, limit int) error {
	query := `insert into manabie_task_limit_setting
				select * from (select ? as user_id, ? as task_id, ? as num_limit, ? as updated_at) as tmp
				where exists( select id from manabie_user where id = ?) and exists(select id from manabie_task where id = ?)
				on duplicate key update num_limit = ?, updated_at = ?`

	now := utils.Now()
	res, err := s.db.Exec(query, userId, taskId, limit, now, userId, taskId, limit, now)
	if err != nil {
		return errors.Wrapf(err, "fail to execute set task limitation")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "fail to get row affected after set task limitation")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(err, fmt.Sprintf("user %d or task %d not exist", userId, taskId))
	}

	return nil
}

func (s *storeImpl) GetAllTaskLimitSetting() ([]togo.UserTaskLimitEntity, error) {
	res := make([]togo.UserTaskLimitEntity, 0)
	query := `select user_id, task_id, num_limit from manabie_task_limit_setting`
	rows, err := s.db.Query(query)
	if err != nil {
		return res, errors.Wrapf(err, "fail to query get all task setting")
	}

	for rows.Next() {
		var r togo.UserTaskLimitEntity
		err = rows.Scan(&r.UserId, &r.TaskId, &r.Limit)
		if err != nil {
			return res, errors.Wrapf(err, "fail to scan task setting row")
		}
		res = append(res, r)
	}

	return res, nil
}
