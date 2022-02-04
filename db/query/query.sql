-- name: GetUserByName :one
SELECT * FROM users WHERE full_name = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (full_name, maximum) VALUES ($1, $2) RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (title, content, is_complete, user_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CountTasks :one
SELECT count(*) FROM tasks WHERE user_id = $1;