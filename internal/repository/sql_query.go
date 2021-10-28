package repository

const (
	SQL_CREATE_USER_TABLE = `
		CREATE TABLE IF NOT EXISTS user (
			id TEXT NOT NULL,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			max_todo INTEGER DEFAULT 5 NOT NULL,
			CONSTRAINT user_PK PRIMARY KEY (id),
			CONSTRAINT user_UK UNIQUE (username)
		);
	`

	SQL_CREATE_TASK_TABLE = `
		CREATE TABLE IF NOT EXISTS task (
			id TEXT NOT NULL,
			content TEXT NOT NULL,
			user_id TEXT NOT NULL,
			created_date TEXT NOT NULL,
			CONSTRAINT task_PK PRIMARY KEY (id),
			CONSTRAINT task_FK FOREIGN KEY (user_id) REFERENCES user(id)
		);
	`

	SQL_TASK_ADD_TASK = `
		INSERT INTO task
		(id, content, user_id, created_date)
		VALUES
		(?, ?, ?, ?)
	`

	SQL_TASK_GET_TASKS = `
		SELECT
			t.id,
			t.content,
			u.username,
			t.created_date
		FROM task t
		INNER JOIN user u ON
			u.id = t.user_id
		WHERE t.user_id = ? AND t.created_date = ?
	`

	SQL_TASK_GET_USER_ID = `
		SELECT id
		FROM user
		WHERE username = ? AND password = ?
		LIMIT 1
	`
)
