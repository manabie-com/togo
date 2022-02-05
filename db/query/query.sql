-- name: GetUserByName :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, daily_task_limit) VALUES ($1, $2) RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (title, content, is_complete, user_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CountUserTasks :one
SELECT count(*) FROM tasks WHERE user_id = $1 AND created_at BETWEEN NOW() - INTERVAL '24 HOURS' AND NOW();
