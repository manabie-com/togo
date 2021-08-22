-- name: ValidateUser :one
SELECT id FROM users WHERE id = $1 AND password = $2;

-- name: CountTaskPerDay :one
SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2;

-- name: GetLimitPerUser :one
SELECT max_todo FROM users WHERE id = $1;

-- name: RetrieveTasks :many
SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2;

-- name: AddTask :exec
INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4);