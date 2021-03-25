-- query.sql

-- name: RetrieveTasks :many
SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2;

-- name: AddTask :one
INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: ValidateUser :one
SELECT id FROM users WHERE id = $1 AND password = $2;

-- name: CountTaskPerDay :one
SELECT count(user_id) FROM tasks WHERE user_id = $1 AND created_date = $2;