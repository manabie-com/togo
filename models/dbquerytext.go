package models

const ( // query for task
	QueryAllTaskText  = "SELECT * FROM tasks WHERE userid = $1"
	FindTaskByIDText  = "SELECT * FROM tasks WHERE id = $1 AND userid = $2"
	InsertTaskText    = "INSERT INTO tasks(content, status, time, timedone, userid) VALUES ($1, $2, $3, $4, $5)"
	DeleteTaskText    = "DELETE FROM tasks WHERE id = $1 AND userid = $2"
	DeleteAllTaskText = "DELETE FROM tasks WHERE userid = $1"
	UpdateTaskText    = "UPDATE tasks SET content =COALESCE($1, content), status = COALESCE($2, status), timedone = COALESCE($3, timedone) WHERE id = $4 AND userid = $5"
)

const ( // query for user
	QueryAllUserText     = "SELECT * FROM users"
	QueryAllUsernameText = "SELECT * FROM users WHERE username = $1"
	FindUserByIDText     = "SELECT * FROM users WHERE id = $1"
	InsertUserText       = "INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3)"
	DeleteUserText       = "DELETE FROM users WHERE id = $1"
	UpdateUserText       = "UPDATE users SET username = COALESCE($1, username), password = COALESCE($2, password), limittask = COALESCE($3, limittask) WHERE id = $4"
)
